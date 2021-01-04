package account

// Account defines an account, i.e. money ledger within Kasikorn
// One customer might have multiple accounts, like a saving account, current account, etc.
type Account struct {
	// Number is the account number in the format XXX-X-XXXXX-X. This is the account number a Kasikorn customer sees
	Number AccountNumber `json:"accountNumber"`

	// Internally ID used by Kasikorn. It is a 14 digits number.
	ID AccountID `json:"id"`

	// Friendly name of the account
	Name string `json:"name"`

	// Nickname, this is more informative than for processing purposes.
	NickName string `json:"nickName"`

	// Type of the account, such as Savings account
	Type AccountType `json:"type"`
}

type AccountNumber string
type AccountID string
