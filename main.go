package main

import (
	"fmt"
	"time"

	"github.com/pinpt/worker/pkg/util"
)

const (
	SignalTimeUnit_NOW          int32 = 0
	SignalTimeUnit_MONTH        int32 = 30
	SignalTimeUnit_QUARTER      int32 = 90
	SignalTimeUnit_HALFYEAR     int32 = 180
	SignalTimeUnit_YEAR         int32 = 365
	SignalTimeUnit_THIRDQUARTER int32 = 270
	SignalTimeUnit_ALLTIME      int32 = -1
	SignalTimeUnit_BIMONTH      int32 = 60
)

func main() {
	// Current functionality

	endDate := GetSignalTime(SignalTimeUnit_NOW, time.Now())
	dateRef := time.Now().UTC()

	date := util.ShortDateFromTime(dateRef)

	startDate := GetSignalTime(30, endDate)

	fmt.Println("StartDate ", startDate)
	fmt.Println("endDate ", endDate)
	fmt.Println("date for Signal table", date)
	// OUTPUT
	// StartDate 2018-03-12 00:00:00 +0000 UTC
	// endDate   2018-04-10 23:59:59.999999999 +0000 UTC
	// date for Signal table 2018-04-11

	// In order to use them in the query we get
	sd := startDate.Unix()
	ed := endDate.Unix()
	fmt.Println("Unix Start Date", sd)
	fmt.Println("Unix End Date", ed)
	// OUTPUT
	// Unix Start Date 1520812800
	// Unix End Date 1523404799

	// In SQL we are saving dates in milliseconds
	// Lets say there are two dates as
	// 1.-
	// 2.-

}

func GetSignalTime(timeUnit int32, refDate time.Time) time.Time {
	var t time.Time
	switch timeUnit {
	case SignalTimeUnit_NOW:
		{
			return refDate.UTC().Truncate(time.Hour * 24)
		}
	case SignalTimeUnit_MONTH:
		{
			t = refDate.UTC().AddDate(0, 0, -29)
		}
	case SignalTimeUnit_BIMONTH:
		{
			t = refDate.UTC().AddDate(0, 0, -59)
		}
	case SignalTimeUnit_QUARTER:
		{
			t = refDate.UTC().AddDate(0, 0, -89)
		}
	case SignalTimeUnit_HALFYEAR:
		{
			t = refDate.UTC().AddDate(0, 0, -179)
		}
	case SignalTimeUnit_THIRDQUARTER:
		{
			t = refDate.UTC().AddDate(0, 0, -269)
		}
	case SignalTimeUnit_YEAR:
		{
			t = refDate.UTC().AddDate(0, 0, -364)
		}
	}

	return t.Truncate(time.Hour * 24)
}

// ShortDateFromTime will return a short date from a time
func ShortDateFromTime(tv time.Time) string {
	return tv.UTC().Format("2006-01-02")
}
