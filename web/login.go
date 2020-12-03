package web

import (
	"io/ioutil"
	"log"
	"net/http"

	"dev.azure.com/noon-homa/Kasikorn/_git/kasikorn/web/login"
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
	log.Printf("Found txtParam: %s", txtParam)
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
	logAccounts(session.accounts)

	errCheckingSession := checkSession(session.cookieJar)
	if errCheckingSession != nil {
		return errCheckingSession
	}
	log.Println("ession worked")

	//
	//errVisintgHome := visitHome(session.cookieJar)
	//if errVisintgHome != nil {
	//	log.Println("Vising home FAILED")
	//} else {
	//	log.Println("Visiting home worked")
	//}

	//token, errStrangeSecWel := login.DoStrangeSecurityWelcome(txtParam, session.cookieJar)
	//if errStrangeSecWel != nil {
	//	return errStrangeSecWel
	//}
	return nil
}

func visitHome(cookies http.CookieJar) error {
	req, errCreatiReq := http.NewRequest(http.MethodGet, kasikornURLs.Home, nil)
	if errCreatiReq != nil {
		return errCreatiReq
	}
	httpClient := http.Client{
		Jar: cookies,
	}
	resp, errDoingReq := httpClient.Do(req)
	if errDoingReq != nil {
		return errDoingReq
	}
	payload, errReadingPayload := ioutil.ReadAll(resp.Body)
	if errReadingPayload != nil {
		return errReadingPayload
	}
	payloadAsString := string(payload)
	log.Printf("Visited home, %d bytes", len(payloadAsString))
	//tokenCandidate := findToken(payloadAsString)
	//if tokenCandidate != nil {
	//	log.Printf("Found token visiting home: %s", *tokenCandidate)
	//}
	return nil
}
