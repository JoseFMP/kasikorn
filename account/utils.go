package account

import (
	"fmt"
	"log"
)

func LogAccounts(accountsToLog map[AccountNumber]*Account) {
	finalMessage := ""
	for _, a := range accountsToLog {
		messageThisAccount := fmt.Sprintf("ID: %s, number: %s", a.ID, a.Number)
		finalMessage = fmt.Sprintf("%s\n%s", finalMessage, messageThisAccount)
	}
	log.Println(finalMessage)
}
