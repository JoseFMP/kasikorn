package account

type AccountType string
type AllAccountTypes struct {
	SavingsAccount AccountType
	CurrentAccount AccountType
}

func GetAllAccountTypes() AllAccountTypes {
	return AllAccountTypes{
		SavingsAccount: "Savings Account",
		CurrentAccount: "Current Account",
	}
}

func getAllAccountTypesMap() map[string]AccountType {
	return map[string]AccountType{
		string(accountTypes.SavingsAccount): accountTypes.SavingsAccount,
		string(accountTypes.CurrentAccount): accountTypes.CurrentAccount,
	}
}

var accountTypes = GetAllAccountTypes()
var accountTypesMap = getAllAccountTypesMap()

func GetAccountType(accountTypeAsString string) AccountType {
	return accountTypesMap[accountTypeAsString]
}
