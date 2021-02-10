package main

import (
	"encoding/json"
	"log"
	"time"

	"dev.azure.com/noon-homa/Kasikorn/_git/kasikorn/account"
	"dev.azure.com/noon-homa/Kasikorn/_git/kasikorn/statement"
	"dev.azure.com/noon-homa/Kasikorn/_git/kasikorn/transaction"
	"dev.azure.com/noon-homa/Kasikorn/_git/kasikorn/web/utils"

	"dev.azure.com/noon-homa/Kasikorn/_git/kasikorn/config"
	"dev.azure.com/noon-homa/Kasikorn/_git/kasikorn/web"
)

// Fill in!
const accountNumber = ``
const username = ""
const password = ""

func main() {

	savingsAccountNumber := account.AccountNumber(accountNumber)
	configuration := config.Config{
		UserName: username,
		Password: password,
	}

	if configuration.UserName == "" || configuration.Password == "" || string(savingsAccountNumber) == "" {
		panic("Fill in user name, password and account number to test")
	}

	session := web.InitSession(configuration)

	for {
		errLogging := session.Login()
		if errLogging != nil {
			log.Println("Retrying...")
			time.Sleep(time.Second * 3)
			continue
		} else {
			break
		}
	}
	log.Println("Logged in")
	from, to := getFromTo()
	statement, errGettingStatement := session.GetStatement(from, to, savingsAccountNumber)

	if errGettingStatement != nil {
		log.Printf("Error getting statement: %v", errGettingStatement)
	} else {
		log.Printf("Statement:\n%+v", *statement)
	}

}

func getFromTo() (utils.KasikornDate, utils.KasikornDate) {

	from := utils.KasikornDate{
		Year: 2020,
		Day:  360,
	}
	to := utils.KasikornDate{
		Year: 2021,
		Day:  11,
	}

	return from, to
}

