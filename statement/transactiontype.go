package statement

import "strings"

type TransactionType string
type AllTransactionTypes struct {
	CashDeposit          TransactionType
	ChequeMoneyTransfer  TransactionType
	TransferWithdrawal   TransactionType
	Comission            TransactionType
	LetterCredit         TransactionType
	InterestDeposit      TransactionType
	WitholdingTaxPayable TransactionType
}

func GetAllTransactionTypes() AllTransactionTypes {
	return AllTransactionTypes{
		CashDeposit:          "Cash Deposit",
		ChequeMoneyTransfer:  "Cheque/Money Transfer",
		TransferWithdrawal:   "Transfer Withdrawal",
		Comission:            "Comission",
		LetterCredit:         "Letter of Credit",
		InterestDeposit:      "Interest Deposit",
		WitholdingTaxPayable: "Witholding Tax Payable",
	}
}

func GetAllTransactionTypesMap() map[string]TransactionType {
	allTransactionTypes := GetAllTransactionTypes()
	return map[string]TransactionType{
		string(allTransactionTypes.CashDeposit):         (allTransactionTypes.CashDeposit),
		string(allTransactionTypes.ChequeMoneyTransfer): allTransactionTypes.ChequeMoneyTransfer,
		string(allTransactionTypes.TransferWithdrawal):  allTransactionTypes.TransferWithdrawal,
	}
}

var allTransactionTypesMap = GetAllTransactionTypesMap()

func transactionTypeFromString(asString string) *TransactionType {
	asString = strings.ReplaceAll(asString, " NB", "")
	//asString = strings.ReplaceAll(asString, " ", "")

	tt, exists := allTransactionTypesMap[asString]
	if !exists {
		return nil
	}
	return &tt
}
