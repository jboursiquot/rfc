# RFC

Side project to understand the history of RFCs published at http://www.ietf.org/download/rfc-index.txt.

### Going from a text document like this

```
0001 Host Software. S. Crocker. April 1969. (Format: TXT=21088 bytes)
     (Status: UNKNOWN) (DOI: 10.17487/RFC0001)

0024 Documentation Conventions. J.F. Heafner, S.D. Crocker. November 1969. (Format:
     TXT=3460 bytes) (Obsoletes RFC0016) (Updates RFC0010, RFC0016)
     (Updated by RFC0027, RFC0030) (Status: UNKNOWN) (DOI:
     10.17487/RFC0024)

0025 No High Link Numbers. S.D. Crocker. October 1969. (Format: TXT=479
     bytes) (Status: UNKNOWN) (DOI: 10.17487/RFC0025)

0026 Not Issued.

0027 Documentation Conventions. S.D. Crocker. December 1969. (Format:
     TXT=3661 bytes) (Updates RFC0010, RFC0016, RFC0024) (Updated by
     RFC0030) (Status: UNKNOWN) (DOI: 10.17487/RFC0027)

0028 Time Standards. W.K. English. January 1970. (Format: TXT=557 bytes)
     (Status: UNKNOWN) (DOI: 10.17487/RFC0028)

0029 Response to RFC 28. R.E. Kahn. January 1970. (Format: TXT=790 bytes)
     (Also RFC0028) (Status: UNKNOWN) (DOI: 10.17487/RFC0029)

0030 Documentation Conventions. S.D. Crocker. February 1970. (Format:
     TXT=4041 bytes) (Updates RFC0010, RFC0016, RFC0024, RFC0027)
     (Status: UNKNOWN) (DOI: 10.17487/RFC0030)

0031 Binary Message Forms in Computer. D. Bobrow, W.R. Sutherland.
     February 1968. (Format: TXT=11191 bytes) (Status: UNKNOWN) (DOI:
     10.17487/RFC0031)

0032 Some Thoughts on SRI's Proposed Real Time Clock. J. Cole. February
     1970. (Format: TXT=2216 bytes) (Status: UNKNOWN) (DOI:
     10.17487/RFC0032)

0033 New Host-Host Protocol. S.D. Crocker. February 1970. (Format:
     TXT=44167 bytes) (Obsoletes RFC0011) (Updated by RFC0036, RFC0047)
     (Status: UNKNOWN) (DOI: 10.17487/RFC0033)
```

### To a parsed and structured version

```
{ID=1129, Title=Internet Time Synchronization: The Network Time Protocol, Authors=D.L. Mills, IssueDate=October 1989, Formats=[], Status=INFORMATIONAL}
{ID=0001, Title=Host Software, Authors=S. Crocker, IssueDate=April 1969, Formats=[Format {Extension=TXT, Bytes=21088}], Status=UNKNOWN}
{ID=0024, Title=Documentation Conventions, Authors=J.F. Heafner, S.D. Crocker, IssueDate=November 1969, Formats=[Format {Extension=TXT, Bytes=3460}], Status=UNKNOWN}
{ID=0025, Title=No High Link Numbers, Authors=S.D. Crocker, IssueDate=October 1969, Formats=[Format {Extension=TXT, Bytes=479}], Status=UNKNOWN}
{ID=0027, Title=Documentation Conventions, Authors=S.D. Crocker, IssueDate=December 1969, Formats=[Format {Extension=TXT, Bytes=3661}], Status=UNKNOWN}
{ID=0028, Title=Time Standards, Authors=W.K. English, IssueDate=January 1970, Formats=[Format {Extension=TXT, Bytes=557}], Status=UNKNOWN}
{ID=0029, Title=Response to RFC 28, Authors=R.E. Kahn, IssueDate=January 1970, Formats=[Format {Extension=TXT, Bytes=790}], Status=UNKNOWN}
{ID=0030, Title=Documentation Conventions, Authors=S.D. Crocker, IssueDate=February 1970, Formats=[Format {Extension=TXT, Bytes=4041}], Status=UNKNOWN}
{ID=0031, Title=Binary Message Forms in Computer, Authors=D. Bobrow, W.R. Sutherland, IssueDate=February 1968, Formats=[Format {Extension=TXT, Bytes=11191}], Status=UNKNOWN}
{ID=0032, Title=Some Thoughts on SRI's Proposed Real Time Clock, Authors=J. Cole, IssueDate=February 1970, Formats=[Format {Extension=TXT, Bytes=2216}], Status=UNKNOWN}
{ID=0033, Title=New Host-Host Protocol, Authors=S.D. Crocker, IssueDate=February 1970, Formats=[Format {Extension=TXT, Bytes=44167}], Status=UNKNOWN}
{ID=0034, Title=Some Brief Preliminary Notes on the Augmentation Research Center Clock, Authors=W.K. English, IssueDate=February 1970, Formats=[Format {Extension=TXT, Bytes=2534}], Status=UNKNOWN}
```

More to come. See [TODO](./TODO.md)
