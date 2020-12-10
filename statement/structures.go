package statement

type Statement struct {
	Meta         StatementMeta
	Transactions []Transaction
}
