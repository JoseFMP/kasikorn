package statement

type Statement struct {
	Meta         StatementMeta `json:"meta"`
	Transactions []Transaction `json:"transactions"`
}
