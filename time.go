package utils

import (
	"errors"
	"time"
)

var timeFormats = []string{
	time.RFC3339,
	"2006-01-02T15:04:05.0000000",
	"2006-01-02T15:04:05",
}

func ParseTime(value string) (time.Time, error) {
	var t time.Time
	var err error
	for _, format := range timeFormats {
		t, err = time.Parse(format, value)
		if err == nil {
			return t, nil
		}
	}
	return time.Time{}, errors.New("unable to parse time: " + value)
}
