DateParse CLI
======================

Simple CLI to convert date string into [unix epoch time](https://en.wikipedia.org/wiki/Unix_time).

The program parses date string into `time.Time` using [dateparse](github.com/araddon/dateparse) library and prints as seconds since epoch, milliseconds since epoch or provided [date layout](https://pkg.go.dev/time#Parse). The program also supports conversion of go duration to seconds or milliseconds.

Examples
----------------------

```console
me@ubuntu:~$ dateparse "Mon Jan 2 15:04:05 MST 2006"
1136214245000
me@ubuntu:~$ dateparse "now"
1661767587158
me@ubuntu:~$ dateparse -o seconds "2009-08-12T22:15:09.99Z"
1250115309
me@ubuntu:~$ dateparse -o millis "18 July 1918"
-1623888000000
me@ubuntu:~$ dateparse -o "2006-01-02T15:04:05-0700" "20-JUN-1990 08:03:00"
1990-06-20T08:03:00+0000
```

Installation
----------------------

```console
me@ubuntu:~$ go install github.com/nilesh-akhade/dateparse@latest
go: downloading github.com/nilesh-akhade/dateparse v0.0.0-20220803091428-df00031503b2
me@ubuntu:~$ dateparse "August 29 2022  10:26:57 GMT"
1661768817000
```
