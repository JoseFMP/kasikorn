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
	}
}

func GetAllTransactionTypesMap() map[string]TransactionType {
	allTransactionTypes := GetAllTransactionTypes()
	return map[string]TransactionType{
		sanitizeTransactionTypeString(string(allTransactionTypes.ChequeWithdrawal)):     allTransactionTypes.ChequeWithdrawal,
		sanitizeTransactionTypeString(string(allTransactionTypes.TransferDeposit)):      allTransactionTypes.TransferDeposit,
		sanitizeTransactionTypeString(string(allTransactionTypes.PaymentReceived)):      allTransactionTypes.PaymentReceived,
		sanitizeTransactionTypeString(string(allTransactionTypes.CashDeposit)):          allTransactionTypes.CashDeposit,
		sanitizeTransactionTypeString(string(allTransactionTypes.CashWithdrawal)):       allTransactionTypes.CashWithdrawal,
		sanitizeTransactionTypeString(string(allTransactionTypes.ChequeMoneyTransfer)):  allTransactionTypes.ChequeMoneyTransfer,
		sanitizeTransactionTypeString(string(allTransactionTypes.ClearingCheque)):       allTransactionTypes.ClearingCheque,
		sanitizeTransactionTypeString(string(allTransactionTypes.TransferWithdrawal)):   allTransactionTypes.TransferWithdrawal,
		sanitizeTransactionTypeString(string(allTransactionTypes.Commission)):           allTransactionTypes.Commission,
		sanitizeTransactionTypeString(string(allTransactionTypes.LetterCredit)):         allTransactionTypes.LetterCredit,
		sanitizeTransactionTypeString(string(allTransactionTypes.InterestDeposit)):      allTransactionTypes.InterestDeposit,
		sanitizeTransactionTypeString(string(allTransactionTypes.WitholdingTaxPayable)): allTransactionTypes.WitholdingTaxPayable,
	}
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
