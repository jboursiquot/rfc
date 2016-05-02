package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/metakeule/fmtdate"
)

// RFC is a Request For Comment
type RFC struct {
	ID          string
	Title       string
	Authors     []string
	IssueDate   time.Time
	Formats     []Format
	Status      string
	DOI         string
	Also        []string
	Updates     []string
	UpdatedBy   []string
	Obsoletes   []string
	ObsoletedBy []string
}

func (rfc RFC) String() string {
	return fmt.Sprintf(
		"RFC {ID=%s, Title=%s, Authors=%s, IssueDate=%s, Formats=%v, Status=%s}",
		rfc.ID, rfc.Title, strings.Join(rfc.Authors, ", "), fmtdate.Format("MMMM YYYY", rfc.IssueDate), rfc.Formats, rfc.Status)
}

// Format includes extention and size in bytes
type Format struct {
	Extension string
	Bytes     int64
}

func (f Format) String() string {
	return fmt.Sprintf("Format {Extension=%s, Bytes=%d}", f.Extension, f.Bytes)
}

var (
	foundRFC bool
	lineSet  []string
	rfcList  = []*RFC{}
	numRe    = regexp.MustCompile(numExpr)
)

func main() {
	file, err := os.Open("./data/rfc-index-sample.txt")
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parseLine(strings.TrimSpace(scanner.Text()))
	}

	parseLineSet() // Ensures the last RFC is parsed.

	if err := scanner.Err(); err != nil {
		fmt.Printf("Err: %s", err)
	}

	log.Printf("%d RFCs", len(rfcList))
}

func parseLine(line string) {
	// log.Printf("parseLine(%s)\n", line)

	if isNotIssued(line) {
		return
	}

	if isNewRFC(line) {
		// log.Println("---isNewRFC -> true")

		if isFirstRFC() {
			newLineSet()
		}

		foundRFC = true

		if isLineSetEmpty() {
			// log.Println("---isLineSetEmpty -> true")
			appendLineSet(line)
		} else {
			// log.Println("---isLineSetEmpty -> false")
			parseLineSet()
			newLineSet()
			appendLineSet(line)
		}
	} else {
		// log.Println("---isNewRFC -> false")
		if isLineSetEmpty() {
			// log.Println("---isLineSetEmpty -> true")
			appendLineSet(line)
		} else {
			// log.Println("---isLineSetEmpty -> false")
			appendLineSet(line)
		}
	}
}

func isFirstRFC() bool {
	return len(rfcList) == 0 && !foundRFC
}

func isNewRFC(line string) bool {
	return len(numRe.FindStringSubmatch(line)) != 0
}

func isLineSetEmpty() bool {
	return len(lineSet) == 0
}

func newLineSet() {
	// log.Println("newLineSet() ***************************")
	lineSet = make([]string, 0)
}

func appendLineSet(line string) {
	// log.Printf("---appendLineSet(%s)\n", line)
	lineSet = append(lineSet, line)
}

func parseLineSet() {
	line := strings.Join(lineSet, " ")

	rfc := &RFC{
		ID:          parseID(line),
		DOI:         parseDOI(line),
		Status:      parseStatus(line),
		Also:        parseList(line, "Also", alsoExpr),
		Updates:     parseList(line, "Updates", updatesExpr),
		UpdatedBy:   parseList(line, "Updated by", updatedByExpr),
		Obsoletes:   parseList(line, "Obsoletes", obsoletesExpr),
		ObsoletedBy: parseList(line, "Obsoleted by", obsoletedByExpr),
	}
	rfc.Title, rfc.Authors = parseTitleAndAuthors(line)
	rfc.Formats = parseFormats(line)

	var err error
	if rfc.IssueDate, err = parseIssueDate(line); err != nil {
		log.Printf("Error->%s, Line->%s", err, line)
	}

	rfcList = append(rfcList, rfc)

	log.Println(rfc)
}
