package statement

import "dev.azure.com/noon-homa/Kasikorn/kasikorn.git/transaction"

type Statement struct {
	Meta         StatementMeta             `json:"meta"`
	Transactions []transaction.Transaction `json:"transactions"`
}
