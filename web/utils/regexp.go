package utils

import (
	"fmt"
	"regexp"
)

const regexpSlashedDatePattern = `\d{2}\/\d{2}\/\d{4}`

func GetRegexpSlashedDate() *regexp.Regexp {
	return regexp.MustCompile(regexpSlashedDatePattern)
}

func GetRegexpSlashedDateAndTime() *regexp.Regexp {
	exp := fmt.Sprintf(`%s\s\d{2}:\d{2}:\d{2}`, regexpSlashedDatePattern)
	//regexpSlashedDateAndTimePattern := regexp.QuoteMeta()
	return regexp.MustCompile(exp)
}
