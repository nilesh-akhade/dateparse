package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/araddon/dateparse"
)

func main() {
	var outputFormat string
	flag.StringVar(&outputFormat, "o", "millis", "millis/seconds/layout")
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Print(`dateparse - Program to parse many date-time formats

Examples:
    dateparse "Mon Jan 2 15:04:05 MST 2006"
    dateparse "now"
    dateparse -o seconds "2009-08-12T22:15:09.99Z"
    dateparse -o millis "18 July 1918"
    dateparse -o "2006-01-02T15:04:05-0700" "20-JUN-1990 08:03:00"

`)
		return
	}

	var err error
	var t time.Time

	dateStr := flag.Args()[0]
	if dateStr == "now" {
		t = time.Now().UTC()
	} else {
		// dateStr := "2022-05-12T16:40:32.171Z"
		t, err = dateparse.ParseAny(dateStr)
		if err != nil {
			fatal(err)
		}
	}

	switch outputFormat {
	case "millis":
		fmt.Println(t.UnixMilli())
	case "seconds":
		fmt.Println(t.Unix())
	default:
		fmt.Println(t.Format(outputFormat))
	}
}

func fatal(err error) {
	fmt.Printf("fatal: %s\n", err)
	os.Exit(1)
}
