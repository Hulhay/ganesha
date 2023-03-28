package utils

import (
	"fmt"
	"time"
)

func Now() time.Time {
	loc, _ := LoadTimeLocation()

	return time.Now().In(loc)
}

func LoadTimeLocation() (*time.Location, error) {
	loc, err := time.LoadLocation(TimeZone)
	if err != nil {
		return nil, err
	}

	return loc, nil
}

func GetTimeOfStartDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

func GetTimeOfEndDay(t time.Time) time.Time {
	return GetTimeOfStartDay(t).AddDate(0, 0, 1).Add(-1 * time.Millisecond)
}

func GetTimeFromString(timeStr string) *time.Time {
	if timeStr == "" {
		return nil
	}

	t, err := time.Parse(TimeFormat, timeStr)
	if err != nil {
		return nil
	}

	return &t
}

func GetStringFromTime(t *time.Time, formatStr string) string {
	if t == nil {
		return ""
	}

	return t.Format(formatStr)
}

func ValidateDateRange(startDate, endDate string, maxDayDuration int) error {
	start, err := time.Parse(DateFormat, startDate)
	if err != nil {
		return ErrInvalidDateFormat
	}

	end, err := time.Parse(DateFormat, endDate)
	if err != nil {
		return ErrInvalidDateFormat
	}

	if end.Before(start) {
		return ErrEndDateBeforeStartDate
	}

	dayDuration := end.Sub(start).Hours() / 24
	if dayDuration > float64(maxDayDuration) {
		return fmt.Errorf("maximal duration date is %d days", maxDayDuration)
	}

	return nil
}
