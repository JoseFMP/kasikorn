package web

import (
	"fmt"
	"log"
	"time"

	"dev.azure.com/noon-homa/Kasikorn/_git/kasikorn/statement"
	"dev.azure.com/noon-homa/Kasikorn/_git/kasikorn/web/statements"
	"dev.azure.com/noon-homa/Kasikorn/_git/kasikorn/web/utils"
)

func validateInput(from time.Time, to time.Time) error {

	if from.IsZero() {
		return fmt.Errorf("From not specified, 0 not valid")
	}

	if to.IsZero() {
		return fmt.Errorf("To not specified, 0 not valid")
	}

	if to.Before(from) {
		return fmt.Errorf("To cannot be before 0")
	}

	now := time.Now()
	oldestFrom := utils.GetSixMonthsFromNowCelinig(now)

	if from.Before(oldestFrom) {
		return fmt.Errorf("Can only fetch statements not older than 6 months")
	}

	return nil
}

func (session *Session) GetStatement(from time.Time, to time.Time, accountNumber string) (*statement.Statement, error) {

	errInput := validateInput(from, to)
	if errInput != nil {
		return nil, errInput
	}
	session.tokenLock.Lock()
	if session.token == nil {
		return nil, fmt.Errorf("No token in session")
	}
	targetAccount, accountAvailable := session.accounts[accountNumber]
	if !accountAvailable {
		return nil, fmt.Errorf("Account %s not found", accountNumber)
	}
	newToken, errSelectingAccount := statements.SelectAccountForStatementInquiry(from, to, targetAccount.ID, *session.token, session.cookieJar)

	if errSelectingAccount != nil {
		session.tokenLock.Unlock()
		return nil, errSelectingAccount
	}
	session.token = &newToken
	session.tokenLock.Unlock()

	session.tokenLock.Lock()
	statementPayload, errDownloading := statements.RequestDownload(from, to, *targetAccount, session.cookieJar, *session.token)
	if errDownloading != nil {
		session.tokenLock.Unlock()

		return nil, errDownloading
	}
	session.tokenLock.Unlock()

	log.Printf("Downloaded %d bytes", len(statementPayload))
	statement, errParsing := statement.Parse(statementPayload)
	if errParsing != nil {
		return nil, errParsing
	}

	return statement, nil
}
