package web

import (
	"log"
	"time"

	"dev.azure.com/noon-homa/Kasikorn/_git/kasikorn/account"
	"dev.azure.com/noon-homa/Kasikorn/_git/kasikorn/web/login"
)

const findTextParamMaxRetries = 3
const durationUserClicking = time.Second * 5 // what is a value that works well? Explore

func (session *Session) Login() error {

	tokenId, errDoingPrelogging := login.DoPreloggin(session.cookieJar)

	if errDoingPrelogging != nil {
		return errDoingPrelogging
	}

	time.Sleep(durationUserClicking) // let's pretend we are a user browsing
	errMakingReqLogin := login.DoLogging(session.cookieJar, tokenId, session.userName, session.password)
	if errMakingReqLogin != nil {
		return errMakingReqLogin
	}

	var txtParam string
	var errFindingTxtParam error
	retries := 0
	for {
		retries++
		log.Println("[Login] Requesting TXT param")
		time.Sleep(durationUserClicking) // let's pretend we are a user browsing
		txtParam, errFindingTxtParam = login.RequestTxtParam(session.cookieJar)
		if errFindingTxtParam == nil {
			break
		}
		if retries > findTextParamMaxRetries {
			return errFindingTxtParam
		}
	}
	log.Printf("[Login] Found txtParam: %d chars", len(txtParam))
	time.Sleep(durationUserClicking) // let's pretend we are a user browsing
	session.tokenLock.Lock()
	token, errDoingSecurityWelcome := login.DoSecurityWelcome(txtParam, session.cookieJar)
	if errDoingSecurityWelcome != nil {
		return errDoingSecurityWelcome
	}
	session.token = &token
	session.tokenLock.Unlock()
	log.Printf("[Login] DoSecurityWelcome found Token: %s", token)

	time.Sleep(durationUserClicking) // let's pretend we are a user browsing
	errRefreshing := session.refreshAccountIDs()
	if errRefreshing != nil {
		return errRefreshing
	}
	accountsLog := account.GenerateLogAccountNumbers(session.accounts)
	log.Printf("[Login] Found accounts: %s", accountsLog)
	errCheckingSession := checkSession(session.cookieJar)
	if errCheckingSession != nil {
		return errCheckingSession
	}
	return nil
}
