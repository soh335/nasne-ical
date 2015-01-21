# nasne-ical

[![wercker status](https://app.wercker.com/status/23363eaa8f758f221810e90ea447b40f/s/master "wercker status")](https://app.wercker.com/project/bykey/23363eaa8f758f221810e90ea447b40f)

convert record schedule of nasne to iCal

## Install

```
$ go get github.com/soh335/nasne-ical
```

## Usage

```
$ nasne-ical --host XXX.XXX.XXX.XXX --name "nasne ical" --tz Asia/Tokyo

BEGIN:VCALENDAR
PRODID:nasne-ical
CALSCALE:GREGORIAN
VERSION:2.0
X-WR-CALNAME:nasne ical
X-WR-CALDESC:nasne ical
X-WR-TIMEZONE:Asia/Tokyo
BEGIN:VEVENT
DTSTAMP:20XXXXXXTXXXXXXZ
UID:XXXXXXXXXXXXXXXXXXXX
TZID:Asia/Tokyo
SUMMARY:XXXXXXXX
DTSTART;TZID=Asia/Tokyo:XXXXXXXXTXXXXXX
DTEND;TZID=Asia/Tokyo:XXXXXXXXTXXXXXX
END:VEVENT
END:VCALENDAR
```
