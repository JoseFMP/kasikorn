package account

// Account defines an account, i.e. money ledger within Kasikorn
// One customer might have multiple accounts, like a saving account, current account, etc.
type Account struct {
	Number   string
	ID       string
	Name     string
	NickName string
	Type     AccountType
}
