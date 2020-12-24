package statement

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"dev.azure.com/noon-homa/Kasikorn/kasikorn.git/web/utils"
)

func parseBody(records [][]string, hasCheckNumber bool) []Transaction {

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

type Transaction struct {
	Date                  time.Time       `json:"date"`
	Type                  TransactionType `json:"type"`
	AmountTHB             float64         `json:"amountTHB"`
	OutstandingBalanceTHB float64         `json:"outstandingBalanceTHB"`
	Channel               ServiceChannel  `json:"channel"`
	Note                  string          `json:"note"`
	CheckNumber           *string         `json:"checkNumber"`
}

type ServiceChannel string
type TransactionType string
type AllTransactionTypes struct {
	CashDepositNB         TransactionType
	ChequeMoneyTransferNB TransactionType
	TransferWithdrawalNB  TransactionType
}

func GetAllTransactionTypes() AllTransactionTypes {
	return AllTransactionTypes{
		CashDepositNB:         "Cash Deposit NB",
		ChequeMoneyTransferNB: "Cheque/Money Transfer NB",
		TransferWithdrawalNB:  "Transfer Withdrawal NB",
	}
}

func GetAllTransactionTypesMap() map[string]TransactionType {
	allTransactionTypes := GetAllTransactionTypes()
	return map[string]TransactionType{
		string(allTransactionTypes.CashDepositNB):         (allTransactionTypes.CashDepositNB),
		string(allTransactionTypes.ChequeMoneyTransferNB): allTransactionTypes.ChequeMoneyTransferNB,
		string(allTransactionTypes.TransferWithdrawalNB):  allTransactionTypes.TransferWithdrawalNB,
	}
}

var allTransactionTypesMap = GetAllTransactionTypesMap()

func transactionTypeFromString(asString string) *TransactionType {
	tt, exists := allTransactionTypesMap[asString]
	if !exists {
		return nil
	}
	return &tt
}

func parseKasikornAmount(amountAsString string) (float64, error) {

	cleanedAmount := strings.ReplaceAll(amountAsString, ",", "")

	amount, errParsing := strconv.ParseFloat(cleanedAmount, 64)
	if errParsing != nil {
		return 0, errParsing
	}
	return amount, nil
}
