package statement

import (
	"fmt"

	"dev.azure.com/noon-homa/Kasikorn/_git/kasikorn/web/utils"
)

func parseBody(records [][]string) {

}

var regexpSlashedDateAndTime = utils.GetRegexpSlashedDateAndTime()

func parseTransaction(record []string) (*Transaction, error) {

	if len(record) < 1 {
		return nil, fmt.Errorf("Record is empty, cannot parse a transaction")
	}

	dateAsString := regexpSlashedDateAndTime.FindString(record[0])
	if dateAsString == "" {
		return nil, fmt.Errorf("Could not find date of this transaction in record")
	}
	transactionTime, errParsing := utils.PlainDateTimeStringToThaiTime(dateAsString)
	if errParsing != nil {
		return nil, errParsing
	}

	return &Transaction{
		Date: *transactionTime,
	}, nil
}
