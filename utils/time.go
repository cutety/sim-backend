package utils

import (
	"sim-backend/utils/logger"
	"time"
)

const GMTISO8601DateFormat = "2006-01-02T15:04:05Z0700"
const DateFormat = "2006-01-02 15:04:05"

func FormatDate(timeString string) *time.Time {
	if timeString == "" {
		return nil
	}
	t, err := time.Parse(DateFormat, timeString)
	if err != nil {
		logger.Info("error:", err)
		return nil
	}
	return &t
}

func ParseWithLocation(timeStr string)  time.Time {
	locationName := "Asia/Shanghai"
	if l, err := time.LoadLocation(locationName); err != nil {
		println(err.Error())
		return time.Time{}
	} else {
		lt, _ := time.ParseInLocation(DateFormat, timeStr, l)
		return lt
	}
}

