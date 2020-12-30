package statement

import "dev.azure.com/noon-homa/Kasikorn/_git/kasikorn/transaction"

type Statement struct {
	Meta         StatementMeta             `json:"meta"`
	Transactions []transaction.Transaction `json:"transactions"`
}
