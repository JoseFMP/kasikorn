package account

import (
	"fmt"
	"log"
	"strings"
)

func LogAccounts(accountsToLog map[AccountNumber]*Account) {
	finalMessage := ""
	for _, a := range accountsToLog {
		messageThisAccount := fmt.Sprintf("ID: %s, number: %s", a.ID, a.Number)
		finalMessage = fmt.Sprintf("%s\n%s", finalMessage, messageThisAccount)
	}
	log.Println(finalMessage)
}

func GenerateLogAccountNumbers(accountsToLog map[AccountNumber]*Account) string {
	accountNumbers := make([]string, len(accountsToLog))
	i := 0
	for _, acc := range accountsToLog {
		accountNumbers[i] = string(acc.Number)
		i++
	}

	res := strings.Join(accountNumbers, ", ")
	return res
}
