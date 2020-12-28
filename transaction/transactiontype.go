package transaction

import "strings"

type TransactionType string
type AllTransactionTypes struct {
	CashWithdrawal       TransactionType
	ClearingCheque       TransactionType
	CashDeposit          TransactionType
	ChequeMoneyTransfer  TransactionType
	TransferWithdrawal   TransactionType
	Commission           TransactionType
	LetterCredit         TransactionType
	InterestDeposit      TransactionType
	WitholdingTaxPayable TransactionType
}

func GetAllTransactionTypes() AllTransactionTypes {
	return AllTransactionTypes{

		CashDeposit:          "Cash Deposit",
		CashWithdrawal:       "Cash Withdrawal",
		ChequeMoneyTransfer:  "Cheque/Money Transfer",
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
