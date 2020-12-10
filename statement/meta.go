package statement

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"dev.azure.com/noon-homa/Kasikorn/_git/kasikorn/account"
	"dev.azure.com/noon-homa/Kasikorn/_git/kasikorn/web/utils"
)

const accountNumberPattern = `\d{3}-\d{1}-\d{5}-\d{1}`

var accountNumberRegExp = regexp.MustCompile(accountNumberPattern)

func parseRecordsMeta(records [][]string) (*StatementMeta, error) {

	result := StatementMeta{
		Account: account.Account{},
	}

	for i, record := range records {

		if i > 6 {
			break
		}
		firstColumn := record[0]

		if strings.Contains(firstColumn, rowHeaders.TransactionHistory) {
			asOfAsString := utils.GetRegexpSlashedDate().FindString(firstColumn)
			asOf, errParsing := utils.PlainDateStringToThaiTime(asOfAsString)
			if errParsing != nil {
				return nil, errParsing
			}
			result.AsOf = asOf
		} else {

			switch firstColumn {
			case rowHeaders.AccountNumber:
				accountNumber := accountNumberRegExp.FindString(record[1])
				if accountNumber == "" {
					return nil, fmt.Errorf("Account number could not be parsed")
				}
				result.Account.Number = accountNumber
			case rowHeaders.AccountName:
				result.Account.Name = record[1]
			case rowHeaders.AccountNickName:
				result.Account.NickName = record[1]
			case rowHeaders.Type:
				accountType := account.GetAccountType(record[1])
				if accountType == "" {
					return nil, fmt.Errorf("Could not parse account type")
				}
				result.Account.Type = accountType
			case rowHeaders.StatementPeriod:
				dates := strings.Split(record[1], "-")
				fromAsString := dates[0]
				toAsString := dates[1]

				from, errParsingFrom := utils.PlainDateStringToThaiTime(fromAsString)
				if errParsingFrom != nil {
					log.Printf("Error parsing from: %s", errParsingFrom)
				}
				to, errParsingTo := utils.PlainDateStringToThaiTime(toAsString)
				if errParsingTo != nil {
					log.Printf("Error parsing to: %s", errParsingTo)
				}
				result.Period.From = from
				result.Period.To = to
			}

		}

	}

	return &result, nil
}

type StatementMeta struct {
	DownloadedOn time.Time
	AsOf         time.Time
	Account      account.Account
	Period       struct {
		From time.Time
		To   time.Time
	}
}
