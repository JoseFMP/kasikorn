package account

import (
	"regexp"
)

const AccountIDPattern = `\d{14}`
const AccountNumberPattern = `\d{3}-\d{1}-\d{5}-\d{1}`

func GetAccountIDRegexp() *regexp.Regexp {
	return regexp.MustCompile(AccountIDPattern)
}

func GetAccountNumberRegexp() *regexp.Regexp {
	return regexp.MustCompile(AccountNumberPattern)
}
