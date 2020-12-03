package utils

import (
	"fmt"
	"time"
)

const dayDuration = time.Hour * 24
const monthDuration = 31 * dayDuration
const sixMonthsCeiling = monthDuration * 6

func GetSixMonthsFromNowCelinig(now time.Time) time.Time {
	return now.Add(-sixMonthsCeiling)
}

const ianaThailandDatabaseTimeZone = "Asia/Bangkok"

func GetThailandTimeZone() *time.Location {
	location := time.FixedZone(ianaThailandDatabaseTimeZone, 7*60*60) // UTC + 7
	return location
}

func PlainDateStringToThaiTime(plainDate string) (*time.Time, error) {
	return parsePlainStringThaiWorld(plainDate, "02/01/2006")
}

func PlainDateTimeStringToThaiTime(plainDate string) (*time.Time, error) {
	return parsePlainStringThaiWorld(plainDate, "02/01/2006 15:04:05")

}

func parsePlainStringThaiWorld(plainDate string, layout string) (*time.Time, error) {
	layoutToUse := fmt.Sprintf("%s +0700", layout)
	result, errParsing := time.Parse(layout, layoutToUse)
	if errParsing != nil {
		return nil, errParsing
	}
	result = result.In(GetThailandTimeZone())
	return &result, nil
}
