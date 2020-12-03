package account

type AccountType string
type AllAccountTypes struct {
	SavingsAccount AccountType
}

func GetAllAccountTypes() AllAccountTypes {
	return AllAccountTypes{
		SavingsAccount: "Savings Account",
	}
}

func getAllAccountTypesMap() map[string]AccountType {
	return map[string]AccountType{
		string(accountTypes.SavingsAccount): accountTypes.SavingsAccount,
	}
}

var accountTypes = GetAllAccountTypes()
var accountTypesMap = getAllAccountTypesMap()

func GetAccountType(accountTypeAsString string) AccountType {
	return accountTypesMap[accountTypeAsString]
}
