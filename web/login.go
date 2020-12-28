package web

import (
	"log"

	"dev.azure.com/noon-homa/Kasikorn/kasikorn.git/account"
	"dev.azure.com/noon-homa/Kasikorn/kasikorn.git/web/login"
)

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
	for {
		txtParam, errFindingTxtParam = login.RequestTxtParam(session.cookieJar)
		if errFindingTxtParam == nil {
			break
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
