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
		sanitizeAccountTypeAsString(string(accountTypes.Savings)): accountTypes.Savings,
		sanitizeAccountTypeAsString(string(accountTypes.Current)): accountTypes.Current,
		sanitizeAccountTypeAsString(string(accountTypes.Deposit)): accountTypes.Deposit,
	}
}

var accountTypes = GetAllAccountTypes()
var accountTypesMap = getAllAccountTypesMap()

func GetAccountType(accountTypeAsString string) AccountType {
	return accountTypesMap[sanitizeAccountTypeAsString(accountTypeAsString)]
}

func sanitizeAccountTypeAsString(asString string) string {
	cleanedAccountType := strings.ReplaceAll(asString, "Account", "")
	cleanedAccountType = strings.ReplaceAll(cleanedAccountType, " ", "")
	return cleanedAccountType
}
