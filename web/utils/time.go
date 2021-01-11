package utils

import (
	"fmt"
	"strings"
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

func PlainDateStringToThaiTime(plainDate string) (time.Time, error) {
	return parsePlainStringThaiWorld(plainDate, "02/01/2006")
}

func PlainDateTimeStringToThaiTime(plainDate string) (time.Time, error) {
	return parsePlainStringThaiWorld(plainDate, "02/01/2006 15:04:05")

}

func parsePlainStringThaiWorld(plainDate string, layout string) (time.Time, error) {

	plainDateCleaned := strings.Trim(plainDate, " ")
	layoutToUse := fmt.Sprintf("%s -0700", layout)
	timeToUse := fmt.Sprintf("%s +0700", plainDateCleaned)

	result, errParsing := time.Parse(layoutToUse, timeToUse)
	if errParsing != nil {
		return time.Time{}, errParsing
	}
	result = result.In(GetThailandTimeZone())
	return result, nil
}

type Period struct {
	From time.Time
	To   time.Time
}

// KasikornDate represents a calendar day in Thailand, i.e. a date without taking into account the time of the day, and meaning within Bangkok time. 1st of January means Day=1.
type KasikornDate struct {
	Year int `json:"year"`
	Day  int `json:"day"`
}

func (kasikornDate KasikornDate) ToTimeDate() time.Time {
	return time.Date(kasikornDate.Year, 1, kasikornDate.Day, 0, 0, 0, 0, thailandTimeZone).In(thailandTimeZone)
}

func (kasikornDate KasikornDate) ToChunks() (string, string, string) {
	targetCalendarDay := kasikornDate.ToTimeDate()
	day := fmt.Sprintf("%02d", targetCalendarDay.Day())
	month := fmt.Sprintf("%02d", targetCalendarDay.Month())
	year := fmt.Sprintf("%02d", targetCalendarDay.Year())

	return day, month, year
}
