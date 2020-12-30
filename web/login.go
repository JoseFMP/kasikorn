package web

import (
	"log"

	"dev.azure.com/noon-homa/Kasikorn/_git/kasikorn/account"
	"dev.azure.com/noon-homa/Kasikorn/_git/kasikorn/web/login"
)

const findTextParamMaxRetries = 5

func (session *Session) Login() error {

	tokenId, errDoingPrelogging := login.DoPreloggin(session.cookieJar)

	if errDoingPrelogging != nil {
		return errDoingPrelogging
	}

	errMakingReqLogin := login.DoLogging(session.cookieJar, tokenId, session.userName, session.password)
	if errMakingReqLogin != nil {
		return errMakingReqLogin
	}

	var txtParam string
	var errFindingTxtParam error
	retries := 0
	for {
		retries++
		log.Println("Requesting TXT param")
		txtParam, errFindingTxtParam = login.RequestTxtParam(session.cookieJar)
		if errFindingTxtParam == nil {
			break
		}
		if retries > findTextParamMaxRetries {
			return errFindingTxtParam
		}
	}
	log.Printf("Found txtParam: %d chars", len(txtParam))
	session.tokenLock.Lock()
	token, errDoingSecurityWelcome := login.DoSecurityWelcome(txtParam, session.cookieJar)
	if errDoingSecurityWelcome != nil {
		return errDoingSecurityWelcome
	}
	session.token = &token
	session.tokenLock.Unlock()
	log.Printf("DoSecurityWelcome found Token: %s", token)

	errRefreshing := session.refreshAccountIDs()
	if errRefreshing != nil {
		return errRefreshing
	}
	accountsLog := account.GenerateLogAccountNumbers(session.accounts)
	log.Printf("Found accounts: %s", accountsLog)
	errCheckingSession := checkSession(session.cookieJar)
	if errCheckingSession != nil {
		return errCheckingSession
	}
	return nil
}
