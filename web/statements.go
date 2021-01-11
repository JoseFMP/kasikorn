package web

import (
	"fmt"
	"log"
	"time"

	"dev.azure.com/noon-homa/Kasikorn/_git/kasikorn/account"
	"dev.azure.com/noon-homa/Kasikorn/_git/kasikorn/statement"
	"dev.azure.com/noon-homa/Kasikorn/_git/kasikorn/web/statements"
	"dev.azure.com/noon-homa/Kasikorn/_git/kasikorn/web/utils"
)

func validateInput(from utils.KasikornDate, to utils.KasikornDate, now time.Time) error {

	if from.Year == 0 {
		return fmt.Errorf("From year not specified, 0 not valid")
	}

	if to.Year == 0 {
		return fmt.Errorf("To not specified, 0 not valid")
	}

	if to.Year < from.Year {
		return fmt.Errorf("To cannot be before 0")
	}

	if to.Year == from.Year {
		if to.Day < from.Day {
			return fmt.Errorf("To cannot be before 0")
		}
	}

	oldestFrom := utils.GetSixMonthsFromNowCelinig(now)

	if oldestFrom.Year() > from.Year {
		return fmt.Errorf("Can only fetch statements not older than 6 months")
	} else if oldestFrom.Year() == from.Year {

		if oldestFrom.YearDay() > from.Day {
			return fmt.Errorf("The year you wanna fetch is ok but the day is before the minimum from")
		}
	}

	return nil
}

func (session *Session) GetStatement(from utils.KasikornDate, to utils.KasikornDate, accountNumber account.AccountNumber) (*statement.Statement, error) {

	errInput := validateInput(from, to, time.Now())
	if errInput != nil {
		return nil, errInput
	}

	time.Sleep(durationUserClicking) // let's pretend we are a user browsing
	session.tokenLock.Lock()
	if session.token == nil {
		return nil, fmt.Errorf("No token in session")
	}
	targetAccount, accountAvailable := session.accounts[accountNumber]
	if !accountAvailable {
		return nil, fmt.Errorf("Account %s not found", accountNumber)
	}
	newToken, downloadCommand, errSelectingAccount := statements.SelectAccountForStatementInquiry(from, to, targetAccount.ID, *session.token, session.cookieJar)

	if errSelectingAccount != nil {
		session.tokenLock.Unlock()
		return nil, errSelectingAccount
	}
	session.token = &newToken
	session.tokenLock.Unlock()

	time.Sleep(durationUserClicking) // let's pretend we are a user browsing
	session.tokenLock.Lock()
	accountIdentificator := statements.AccountIdentificator{
		ID:     targetAccount.ID,
		Number: targetAccount.Number,
	}
	statementPayload, errDownloading := statements.RequestDownload(from, to, accountIdentificator, downloadCommand, session.cookieJar, *session.token)
	if errDownloading != nil {
		session.tokenLock.Unlock()

		return nil, errDownloading
	}
	session.tokenLock.Unlock()

	log.Printf("[%s]Downloaded %d bytes", string(targetAccount.Number), len(statementPayload))
	statement, errParsing := statement.Parse(statementPayload)
	if errParsing != nil {
		return nil, errParsing
	}

	return statement, nil
}
