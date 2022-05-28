DateParse CLI
======================

Simple CLI to parse given date into expected format

Examples:

    dateparse "Mon Jan 2 15:04:05 MST 2006"
    dateparse "now"
    dateparse -o seconds "2009-08-12T22:15:09.99Z"
    dateparse -o millis "18 July 1918"
    dateparse -o "2006-01-02T15:04:05-0700" "20-JUN-1990 08:03:00"