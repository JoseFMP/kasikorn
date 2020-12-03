package web

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"

	"dev.azure.com/noon-homa/Kasikorn/_git/kasikorn/account"
	"dev.azure.com/noon-homa/Kasikorn/_git/kasikorn/web/token"
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

const patternRegExp = `<option value="\d{14}">\d{3}-\d{1}-\d{5}-\d{1}\s.*</option>`

func getAccountRegexp() *regexp.Regexp {
	return regexp.MustCompile(patternRegExp)
}

func parseAccountsInPayload(payload string) []*account.Account {
	accounts := make([]*account.Account, 0)
	matches := getAccountRegexp().FindAllString(payload, -1)
	for _, m := range matches {
		candidateAccount := processCandidateAccount(m)
		if candidateAccount != nil {
			accounts = append(accounts, candidateAccount)
		}
	}
	return accounts
}

const accountIdPattern = `value="\d{14}"`
const accountNumberPattern = `>\d{3}-\d{1}-\d{5}-\d{1}`

func getAccountIDRegexp() *regexp.Regexp {
	return regexp.MustCompile(accountIdPattern)
}

func getAccountNumberRegexp() *regexp.Regexp {
	return regexp.MustCompile(accountNumberPattern)
}

func processCandidateAccount(payload string) *account.Account {

	idChunk := getAccountIDRegexp().FindString(payload)
	idChunk = strings.ReplaceAll(idChunk, `value="`, "")
	id := strings.ReplaceAll(idChunk, `"`, "")

	numberChunk := getAccountNumberRegexp().FindString(payload)
	numberChunk = strings.ReplaceAll(numberChunk, `>`, "")
	number := strings.ReplaceAll(numberChunk, ` `, "")

	if id != "" && number != "" {
		return &account.Account{
			ID:     id,
			Number: number,
		}
	}

	return nil
}

func logAccounts(accountsToLog map[string]*account.Account) {
	finalMessage := ""
	for _, a := range accountsToLog {
		messageThisAccount := fmt.Sprintf("ID: %s, number: %s", a.ID, a.Number)
		finalMessage = fmt.Sprintf("%s\n%s", finalMessage, messageThisAccount)
	}
	log.Println(finalMessage)
}
