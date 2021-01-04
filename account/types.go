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

	stringSanitized := sanitizeAccountTypeAsString(accountTypeAsString)
	return accountTypesMap[stringSanitized]
}

func sanitizeAccountTypeAsString(asString string) string {
	cleanedAccountType := strings.ToLower(asString)
	cleanedAccountType = strings.ReplaceAll(cleanedAccountType, "account", "")
	cleanedAccountType = strings.ReplaceAll(cleanedAccountType, " ", "")
	return cleanedAccountType
}
