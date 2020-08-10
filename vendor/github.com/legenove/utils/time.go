/*
	time utils
*/
package utils

import (
	"fmt"
	"time"
)

const (
	TimeFormat      = "2006-01-02 15:04:05"
	DateFormat      = "2006-01-02"
	YearFormat      = "2006"
	DeufautTimeZone = "UTC"
	Time            = "15:04:05"
	TimezoneFormat  = "UTC_%.1f"
)

var (
	TimeZoneMap = make(map[string]*time.Location, 60)
)

func init() {
	var i float64
	for i = -12; i <= 12; i += 0.5 {
		TimeZoneMap[fmt.Sprintf(TimezoneFormat, i)] = getTimeLocal(i)
	}
}

// get default location
func TimeLoadLocation() (*time.Location, error) {
	return time.LoadLocation(DeufautTimeZone)
}

// get UTC time string format by _fmt
func UTCTimeString(t time.Time, _fmt string) string {
	l, _ := TimeLoadLocation()
	return t.In(l).Format(_fmt)
}

// get current time
func GetCurrentTime() time.Time {
	l, _ := TimeLoadLocation()
	return time.Now().In(l)
}

// get current time string by format
func GetCurrentFormat(format string) string {
	return UTCTimeString(GetCurrentTime(), format)
}

// get current year
func GetCurrentYear() string {
	return GetCurrentFormat(YearFormat)
}

// get current date
func GetCurrentDate() string {
	return GetCurrentFormat(DateFormat)
}

// timezone Location
func getTimeLocal(timezone float64) *time.Location {
	return time.FixedZone(fmt.Sprintf(TimezoneFormat, timezone), int(timezone*3600))
}

// convert timeUnix to Time
func GetTimeUnixToTime(timeUnix int64) time.Time {
	return time.Unix(timeUnix, 0)
}

// format time by format and timezone
func FormatTimeByTimezone(t time.Time, format string, timezone float64) string {
	l, ok := TimeZoneMap[fmt.Sprintf(TimezoneFormat, timezone)]
	if !ok {
		l = getTimeLocal(timezone)
		TimeZoneMap[fmt.Sprintf(TimezoneFormat, timezone)] = l
	}
	return t.In(l).Format(format)
}
