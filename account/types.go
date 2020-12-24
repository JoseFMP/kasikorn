package account

import (
	"strings"
)

type AccountType string
type AllAccountTypes struct {
	Savings AccountType
	Current AccountType
	Deposit AccountType
}

func GetAllAccountTypes() AllAccountTypes {
	return AllAccountTypes{
		Savings: "Savings",
		Current: "Current",
		Deposit: "Deposit",
	}
}

func getAllAccountTypesMap() map[string]AccountType {
	return map[string]AccountType{
		string(accountTypes.Savings): accountTypes.Savings,
		string(accountTypes.Current): accountTypes.Current,
	}
}

var accountTypes = GetAllAccountTypes()
var accountTypesMap = getAllAccountTypesMap()

func GetAccountType(accountTypeAsString string) AccountType {

	cleanedAccountType := strings.ReplaceAll(accountTypeAsString, "Account", "")
	cleanedAccountType = strings.ReplaceAll(cleanedAccountType, " ", "")
	return accountTypesMap[cleanedAccountType]
}
