package utils

import "errors"

const (
	TimestamptzFormat string = "2006-01-02 15:04:05 +0000 MST"
	TimeFormat        string = "2006-01-02 15:04:05.000-0700"
	TimeFormatDB      string = "2006-01-02 15:04:05"
	DateFormat        string = "2006-01-02"
	TimeZone          string = "Asia/Jakarta"
)

var (
	ErrInvalidDateFormat      = errors.New("invalid date format")
	ErrEndDateBeforeStartDate = errors.New("end date must be after start date")
)
