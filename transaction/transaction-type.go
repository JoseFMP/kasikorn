package transaction

import "strings"

type TransactionType string
type AllTransactionTypes struct {
	TransferDeposit TransactionType
	InterestDeposit TransactionType
	CashDeposit     TransactionType

	PaymentReceived TransactionType

	CashWithdrawal     TransactionType
	TransferWithdrawal TransactionType
	ChequeWithdrawal   TransactionType

	ClearingCheque      TransactionType
	ChequeMoneyTransfer TransactionType

	Commission   TransactionType
	LetterCredit TransactionType

	WitholdingTaxPayable TransactionType

	OpenAccount TransactionType
}

func GetAllTransactionTypes() AllTransactionTypes {
	return AllTransactionTypes{
		TransferDeposit:      "Transfer Deposit",
		PaymentReceived:      "Payment Received",
		CashDeposit:          "Cash Deposit",
		CashWithdrawal:       "Cash Withdrawal",
		ChequeMoneyTransfer:  "Cheque/Money Transfer",
		ChequeWithdrawal:     "Cheque Withdrawal",
		ClearingCheque:       "Clearing Cheque",
		TransferWithdrawal:   "Transfer Withdrawal",
		Commission:           "Commission",
		LetterCredit:         "Letter of Credit",
		InterestDeposit:      "Interest Deposit",
		WitholdingTaxPayable: "Withholding Tax Payable",
		OpenAccount:          "Open Account",
	}
}

func GetAllTransactionTypesMap() map[string]TransactionType {
	mapToUse := make(map[string]TransactionType)

	addToMap(mapToUse, allTransactionTypes.ChequeWithdrawal)
	addToMap(mapToUse, allTransactionTypes.TransferDeposit)
	addToMap(mapToUse, allTransactionTypes.PaymentReceived)
	addToMap(mapToUse, allTransactionTypes.CashDeposit)
	addToMap(mapToUse, allTransactionTypes.CashWithdrawal)
	addToMap(mapToUse, allTransactionTypes.ChequeMoneyTransfer)
	addToMap(mapToUse, allTransactionTypes.ClearingCheque)
	addToMap(mapToUse, allTransactionTypes.TransferWithdrawal)
	addToMap(mapToUse, allTransactionTypes.Commission)
	addToMap(mapToUse, allTransactionTypes.LetterCredit)
	addToMap(mapToUse, allTransactionTypes.InterestDeposit)
	addToMap(mapToUse, allTransactionTypes.WitholdingTaxPayable)
	addToMap(mapToUse, allTransactionTypes.WitholdingTaxPayable)
	addToMap(mapToUse, allTransactionTypes.OpenAccount)

	return mapToUse
}

func addToMap(mapToUse map[string]TransactionType, transactionType TransactionType) {
	sanitized := sanitizeTransactionTypeString(string(transactionType))

	mapToUse[sanitized] = transactionType
}

var allTransactionTypesMap = GetAllTransactionTypesMap()

func transactionTypeFromString(asString string) *TransactionType {
	sanitizedString := sanitizeTransactionTypeString(asString)
	tt, exists := allTransactionTypesMap[sanitizedString]
	if !exists {
		return nil
	}
	return &tt
}

func sanitizeTransactionTypeString(dangerousString string) string {

	newString := strings.ReplaceAll(dangerousString, " NB", "")
	newString = strings.ReplaceAll(newString, " ", "")
	newString = strings.ToLower(newString)

	return newString
}
