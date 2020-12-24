package main

import (
	"log"
	"time"

	"dev.azure.com/noon-homa/Kasikorn/kasikorn.git/account"
	"dev.azure.com/noon-homa/Kasikorn/kasikorn.git/web/utils"

	"dev.azure.com/noon-homa/Kasikorn/kasikorn.git/config"
	"dev.azure.com/noon-homa/Kasikorn/kasikorn.git/web"
)

func main() {

	// Fill in!
	savingsAccountNumber := account.AccountNumber(``)
	configuration := config.Config{
		UserName: "",
		Password: "",
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
		log.Printf("Error Getting statement")
	} else {
		log.Printf("Statement:\n%+v", *statement)
	}
}

func getFromTo() (time.Time, time.Time) {

	now := time.Now()

	from := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, utils.GetThailandTimeZone())
	to := now.Add(-time.Hour * 24)
	log.Printf("From: %s, to: %s", from.Format("2006/01/02"), to.Format("2006/01/02"))
	return from, to
}
