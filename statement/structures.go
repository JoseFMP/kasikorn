package statement

import "github.com/JoseFMP/kasikorn/transaction"

type Statement struct {
	Meta         StatementMeta             `json:"meta"`
	Transactions []transaction.Transaction `json:"transactions"`
}
