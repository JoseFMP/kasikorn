package login

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"dev.azure.com/noon-homa/Kasikorn/_git/kasikorn/web/token"
)

const patternTxtParam = `<input type="hidden" name="txtParam" value="`

func RequestTxtParam(cookieJar http.CookieJar) (string, error) {
	req, errCreatingReq := http.NewRequest(http.MethodGet, kasikornURLs.RedirectIB, nil)
	if errCreatingReq != nil {
		return "", errCreatingReq
	}
	httpClient := http.Client{
		Jar: cookieJar,
	}
	resp, errDoingR := httpClient.Do(req)
	if errDoingR != nil {
		return "", errDoingR
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP status failed: %s", resp.Status)
	}

	respPayload, errReadingPayload := ioutil.ReadAll(resp.Body)
	if errReadingPayload != nil {
		return "", errReadingPayload
	}
	payloadAsString := string(respPayload)

	targetString := parseTxtParam(payloadAsString)
	tokenCandidate := token.FindToken(payloadAsString)
	if tokenCandidate != nil {
		log.Printf("txtParam found token: %s", *tokenCandidate)
	}
	if targetString == nil {
		return "", fmt.Errorf("Could not parse txtParam from payload")
	}
	return *targetString, nil
}

func parseTxtParam(payload string) *string {
	delimiterStart := strings.Index(payload, patternTxtParam)

	if delimiterStart == -1 {
		return nil
	}

	targetString := payload[(delimiterStart + len(patternTxtParam)):]
	delimiterEnd := strings.Index(targetString, `"`)
	if delimiterEnd == -1 {
		return nil
	}
	targetString = targetString[:delimiterEnd]
	return &targetString
}
