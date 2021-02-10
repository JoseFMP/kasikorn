package web

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"

	"github.com/JoseFMP/kasikorn/account"
	"github.com/JoseFMP/kasikorn/web/token"
)

func (session *Session) refreshAccountIDs() error {
	if session.token == nil {
		return fmt.Errorf("Cannot refresh the accounts as no session token available")
	}
	session.tokenLock.Lock()
	accounts, newToken, errFindingAccounts := findAccounts(session.cookieJar)
	if errFindingAccounts != nil {
		return errFindingAccounts
	}

	session.token = &newToken
	log.Printf("refreshAccountIDs updated token: %s", *session.token)
	session.tokenLock.Unlock()

	for _, a := range accounts {
		session.accounts[a.Number] = a
	}
	return nil
}

func findAccounts(cookies http.CookieJar) ([]*account.Account, string, error) {
	req, errCreatingReq := http.NewRequest(http.MethodGet, eBankURLs.StatementInquiry, nil)

	if errCreatingReq != nil {
		return nil, "", errCreatingReq
	}
	httpClient := http.Client{
		Jar: cookies,
	}
	resp, errDoingReq := httpClient.Do(req)
	if errDoingReq != nil {
		return nil, "", errDoingReq
	}
	if resp.StatusCode != http.StatusOK {
		return nil, "", fmt.Errorf("Error making HTTP request: %s", resp.Status)
	}
	payload, errReadingPayload := ioutil.ReadAll(resp.Body)

	if errReadingPayload != nil {
		return nil, "", errReadingPayload
	}
	payloadAsString := string(payload)
	candidateToken := token.FindToken(payloadAsString)
	if candidateToken != nil {
		log.Printf("Token found in findAccounts: %s", *candidateToken)
	} else {
		return nil, "", fmt.Errorf("Did not find token in findAccounts :(")
	}
	accounts := parseAccountsInPayload(payloadAsString)
	return accounts, *candidateToken, nil
}

func getAccountOptionRegexp() *regexp.Regexp {

	patternRegExp := fmt.Sprintf(`<option value="%s">%s\s.*</option>`, account.AccountIDPattern, account.AccountNumberPattern)
	return regexp.MustCompile(patternRegExp)
}

func parseAccountsInPayload(payload string) []*account.Account {
	accounts := make([]*account.Account, 0)
	matches := getAccountOptionRegexp().FindAllString(payload, -1)
	for _, m := range matches {
		candidateAccount := processCandidateAccount(m)
		if candidateAccount != nil {
			accounts = append(accounts, candidateAccount)
		}
	}
	return accounts
}

func getAccountIDInHTMLRegexp() *regexp.Regexp {
	accountIDInHTMLPatter := fmt.Sprintf(`value="%s"`, account.AccountIDPattern)
	return regexp.MustCompile(accountIDInHTMLPatter)
}

func getAccountNumberInHTMLRegexp() *regexp.Regexp {

	accountNumberInHTMLPattern := fmt.Sprintf(`>%s`, account.AccountNumberPattern)
	return regexp.MustCompile(accountNumberInHTMLPattern)
}

func processCandidateAccount(payload string) *account.Account {

	idChunk := getAccountIDInHTMLRegexp().FindString(payload)
	id := account.GetAccountIDRegexp().FindString(idChunk)

	numberChunk := getAccountNumberInHTMLRegexp().FindString(payload)
	number := account.GetAccountNumberRegexp().FindString(numberChunk)

	if id != "" && number != "" {
		return &account.Account{
			ID:     account.AccountID(id),
			Number: account.AccountNumber(number),
		}
	}

	return nil
}
