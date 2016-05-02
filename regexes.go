package main

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/metakeule/fmtdate"
)

const (
	numExpr          = `^[0-9]{4}\s`
	titleAuthorsExpr = `([\s\S])+`
	issueDateExpr    = `(January|February|March|April|May|June|July|August|September|October|November|December)\s[0-9]{4}`
	notIssuedExpr    = `^[0-9]{4}\sNot Issued.`
	formatExpr       = `(\(Format:\s(TXT|PS|PDF)=([0-9]*)\sbytes\))`
	formatExtExpr    = `(TXT|PS|PDF)`
	formatLenExpr    = `[0-9]{1,}`
	formatExtLenExpr = `(TXT|PS|PDF)=([0-9]*)`
	statusExpr       = `(\(Status:\s([a-zA-Z]*)\))`
	doiExpr          = `(\(DOI:\s([a-zA-Z0-9./]*\)))`
	alsoExpr         = `\((Also\s([(RFC|FYI|STD|BCP)0-9]*))\)`
	updatesExpr      = `\(Updates(\s([RFC0-9],?)*)*\)`
	updatedByExpr    = `\(Updated\sby(\s([RFC0-9],?)*)*\)`
	obsoletesExpr    = `\(Obsoletes(\s([RFC0-9],?)*)*\)`
	obsoletedByExpr  = `\(Obsoleted\sby(\s([RFC0-9],?)*)*\)`
)

const (
	unspecified = "UNSPECIFIED"
)

func isNotIssued(line string) bool {
	return regexp.MustCompile(notIssuedExpr).FindString(line) != ""
}

func parseID(line string) string {
	return strings.TrimSpace(regexp.MustCompile(numExpr).FindString(line))
}

func parseTitleAndAuthors(line string) (string, []string) {
	compExpr := numExpr + titleAuthorsExpr + `\s` + issueDateExpr
	title := regexp.MustCompile(compExpr).FindString(line)
	title = strings.TrimSpace(strings.Replace(title, parseID(line), "", -1))
	title = strings.TrimSpace(strings.Replace(title, parseIssueDateStr(line), "", -1))
	title = strings.TrimSuffix(title, ".")
	authors := strings.Split(title, ", ")
	title = strings.Split(title, ". ")[0]
	authors[0] = strings.Replace(authors[0], title+". ", "", -1)
	return title, authors
}

func parseIssueDateStr(line string) string {
	return regexp.MustCompile(issueDateExpr).FindString(line)
}

func parseIssueDate(line string) (time.Time, error) {
	dateStr := parseIssueDateStr(line)
	return fmtdate.Parse("MMMM YYYY", dateStr)
}

func parseFormats(line string) []Format {
	var str string
	if str = regexp.MustCompile(formatExpr).FindString(line); str == "" {
		return make([]Format, 0)
	}

	matches := regexp.MustCompile(formatExtLenExpr).FindAllString(str, -1)

	var formats []Format
	for _, f := range matches {
		s := strings.Split(f, "=")
		bytes, _ := strconv.ParseInt(s[1], 10, 64)
		formats = append(formats, Format{Extension: s[0], Bytes: bytes})
	}

	return formats
}

func parseDOI(line string) string {
	var str string
	if str = regexp.MustCompile(doiExpr).FindString(line); str == "" {
		return unspecified
	}
	return strings.TrimSuffix(strings.Split(str, ": ")[1], ")")
}

func parseStatus(line string) string {
	var str string
	if str = regexp.MustCompile(statusExpr).FindString(line); str == "" {
		return unspecified
	}
	return strings.TrimSuffix(strings.Split(str, ": ")[1], ")")
}

func parseAlso(line string) string {
	return regexp.MustCompile(alsoExpr).FindString(line)
}

func parseList(line string, name string, expr string) []string {
	var str string
	if str = regexp.MustCompile(expr).FindString(line); str == "" {
		return make([]string, 0)
	}
	str = strings.Replace(str, "("+name+" ", "", -1)
	return strings.Split(strings.TrimSuffix(str, ")"), ", ")
}

func parseObsoletes(line string) string {
	return regexp.MustCompile(obsoletesExpr).FindString(line)
}

func parseObsoletedBy(line string) string {
	return regexp.MustCompile(obsoletedByExpr).FindString(line)
}
