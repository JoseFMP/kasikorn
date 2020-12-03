package statement

import (
	"time"

	"dev.azure.com/noon-homa/Kasikorn/_git/kasikorn/account"
)

type Statement struct {
	Meta    StatementMeta
	Records Transaction
}

type StatementMeta struct {
	DownloadedOn time.Time
	AsOf         time.Time
	Account      account.Account
}

type Transaction struct {
	Date                  time.Time
	TransactionType       TransactionType
	AmountTHB             float64
	OutstandingBalanceTHB float64
	Channel               TransactionChannel
	Note                  string
}

type TransactionChannel string
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
