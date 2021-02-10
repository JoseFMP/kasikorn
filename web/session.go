package web

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"strings"
	"sync"

	"github.com/JoseFMP/kasikorn/account"
	"github.com/JoseFMP/kasikorn/config"
	"github.com/JoseFMP/kasikorn/web/utils"
)

type Session struct {
	userName      string
	password      string
	token         *string
	tokenLock     sync.RWMutex
	cookieJar     http.CookieJar
	accounts      map[account.AccountNumber]*account.Account
	initialConfig *config.Config
}

func InitSession(config config.Config) *Session {

	cookieJarOptions := cookiejar.Options{}
	cookieJar, errCreatingCookieJar := cookiejar.New(&cookieJarOptions)
	if errCreatingCookieJar != nil {
		panic("Could not create cookie jar")
	}

	session := &Session{
		userName:      config.UserName,
		password:      config.Password,
		accounts:      make(map[account.AccountNumber]*account.Account),
		cookieJar:     cookieJar,
		initialConfig: &config,
	}
	return session
}

func checkSession(cookies http.CookieJar) error {
	req, errMakingRe := http.NewRequest(http.MethodPost, kasikornURLs.CheckSession, nil)
	if errMakingRe != nil {
		return errMakingRe
	}
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Accept-Language", "en-US,en;q=0.9,es;q=0.8,fr;q=0.7,th;q=0.6,de;q=0.5")

	req.Header.Add("Host", kasikornURLs.HostName)
	req.Header.Add("Origin", kasikornURLs.Origin)
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-Fetch-Site", "same-origin")

	httpClient := http.Client{
		Jar: cookies,
	}

	resp, errDoingReq := httpClient.Do(req)
	if errDoingReq != nil {
		return errDoingReq
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Not OK HTTP reply...")
	}

	payloadAsBytes, errReadingPayload := ioutil.ReadAll(resp.Body)
	if errReadingPayload != nil {
		return errReadingPayload
	}

	payloadAsString := string(payloadAsBytes)

	if !strings.Contains(payloadAsString, `<result>true</result>`) {
		return fmt.Errorf("Check session failed")
	}

	return nil
}

var eBankURLs = utils.GetAllEbankURLs()
var kasikornURLs = utils.GetAllOnlineKasikornURLs()
