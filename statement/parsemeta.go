package statement

import (
	"fmt"
	"regexp"
	"strings"

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
			result.AsOf = *asOf
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
			}

		}

	}

	return &result, nil
}
