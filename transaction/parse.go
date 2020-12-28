package transaction

import (
	"fmt"
	"log"

	"dev.azure.com/noon-homa/Kasikorn/kasikorn.git/web/utils"
)

func Parse(records [][]string, hasCheckNumber bool) []Transaction {

	transactions := make([]Transaction, 0)

	for _, r := range records {
		transaction, errParsingRecord := parseTransaction(r, hasCheckNumber)
		if errParsingRecord == nil {
			transactions = append(transactions, transaction)
		} else {
			log.Printf("Error parsing transaction: %v", errParsingRecord)
		}
	}
	return transactions
}

var regexpSlashedDateAndTime = utils.GetRegexpSlashedDateAndTime()

func parseTransaction(record []string, hasCheckNumber bool) (Transaction, error) {

	checkNumberOffset := 0
	if hasCheckNumber {
		checkNumberOffset = 1
	}
	if len(record) == 0 {
		return Transaction{}, fmt.Errorf("Record is empty, cannot parse a transaction")
	}

	if (hasCheckNumber && len(record) < 8) || (!hasCheckNumber && len(record) < 7) {
		return Transaction{}, fmt.Errorf("Record does not have enough columns")
	}

	dateAsString := regexpSlashedDateAndTime.FindString(record[0])
	if dateAsString == "" {
		return Transaction{}, fmt.Errorf("Could not find date of this transaction in record")
	}
	transactionTime, errParsing := utils.PlainDateTimeStringToThaiTime(dateAsString)
	if errParsing != nil {
		return Transaction{}, errParsing
	}

	transactionType := transactionTypeFromString(record[1])
	if transactionType == nil {
		return Transaction{}, fmt.Errorf(`Not found transaction type "%s"`, record[1])
	}

	var checkNumber *string
	if hasCheckNumber {
		checkNumber = &record[2]
	}

	amountAsString := record[2+checkNumberOffset]
	deposit := false

	if record[3+checkNumberOffset] != "" {
		deposit = true
		amountAsString = record[3+checkNumberOffset]
	}

	amount, errParsing := parseKasikornAmount(amountAsString)
	if errParsing != nil {
		return Transaction{}, fmt.Errorf("Cannot parse amount of transaction: %s", amountAsString)
	}
	if !deposit {
		amount = amount * (-1)
	}

	balance, errParsingBalance := parseKasikornAmount(record[4+checkNumberOffset])
	if errParsingBalance != nil {
		return Transaction{}, errParsingBalance
	}

	return Transaction{
		Date:                  transactionTime,
		Type:                  *transactionType,
		AmountTHB:             amount,
		OutstandingBalanceTHB: balance,
		Channel:               ServiceChannel(record[5]),
		Note:                  record[6],
		CheckNumber:           checkNumber,
	}, nil
}
