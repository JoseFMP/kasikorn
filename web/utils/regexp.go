package utils

import "regexp"

func GetRegexpSlashedDate() *regexp.Regexp {
	regexpSlashedDatePattern := `\d{2}\/\d{2}\/\d{4}`
	return regexp.MustCompile(regexpSlashedDatePattern)
}

func GetRegexpSlashedDateAndTime() *regexp.Regexp {
	regexpSlashedDateAndTimePattern := `\d{2}\/\d{2}\/\d{4}\s\d{2}:\d{2}:d\{2}`
	return regexp.MustCompile(regexpSlashedDateAndTimePattern)
}
