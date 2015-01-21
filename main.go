package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/soh335/nasne"
	"github.com/soh335/nasne/xsrs"
)

var (
	host   = flag.String("host", "", "host of nasne")
	port   = flag.String("port", "64230", "port of nasne for request xsrs action")
	name   = flag.String("name", "", "name of ical")
	tz     = flag.String("tz", "", "timezone. such as Asia/Tokyo")
	prodid = flag.String("prodid", "nasne-ical", "prodid of ical")
)

func main() {
	flag.Parse()

	if err := _main(); err != nil {
		log.Fatal(err)
	}
}

func _main() error {
	if *host == "" {
		return fmt.Errorf("host is required")
	}

	if *name == "" {
		return fmt.Errorf("name is required")
	}

	if *tz == "" {
		return fmt.Errorf("tz is required")
	}
	r, err := nasne.GetRecordScheduleList(net.JoinHostPort(*host, *port))
	if err != nil {
		return err
	}
	return _ical(os.Stdout, r, *tz, *name)
}

func _ical(w io.Writer, r *xsrs.Root, tz string, name string) error {
	var c VCalendar
	c.PRODID = *prodid
	c.VERSION = "2.0"
	c.CALSCALE = "GREGORIAN"

	c.X_WR_CALNAME = name
	c.X_WR_CALDESC = name
	c.X_WR_TIMEZONE = tz

	components := []VComponent{}
	dtstamp := time.Now()
	layout := "2006-01-02T15:04:05-0700"

	for _, item := range r.Items {
		dtstart, err := time.Parse(layout, item.ScheduledStartDateTime)
		if err != nil {
			return err
		}
		duration, err := strconv.Atoi(item.ScheduledDuration)
		if err != nil {
			return err
		}
		dtend := dtstart.Add(time.Second * time.Duration(duration))

		var e VEvent
		e.UID = item.Id
		e.DTSTAMP = dtstamp
		e.DTSTART = dtstart
		e.DTEND = dtend
		e.SUMMARY = item.Title
		e.TZID = tz

		components = append(components, &e)
	}

	c.VComponent = components
	return c.Encode(w)
}
