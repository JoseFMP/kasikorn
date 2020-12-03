package login

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"dev.azure.com/noon-homa/Kasikorn/_git/kasikorn/web/token"
	"dev.azure.com/noon-homa/Kasikorn/_git/kasikorn/web/utils"
)

func DoSecurityWelcome(txtParam string, cookies http.CookieJar) (string, error) {

	reqPayload := map[string]string{
		"txtParam": txtParam,
	}
	req, errCreatingReq := utils.CreatePostFormReq(reqPayload, eBankURLs.SecurityWelcome)

	if errCreatingReq != nil {
		return "", errCreatingReq
	}

	redirected := false
	httpClient := http.Client{
		Jar: cookies,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			redirected = true
			log.Printf("Secuirty welcome doing redirect to %s", req.URL.String())
			return nil //http.ErrUseLastResponse
		},
	}

	resp, errDoingReq := httpClient.Do(req)
	if errDoingReq != nil {
		return "", errDoingReq
	}

	if redirected {

		payload, errReadingPayload := ioutil.ReadAll(resp.Body)
		if errReadingPayload != nil {
			return "", errReadingPayload
		}
		payloadAsString := string(payload)
		tokenCandidate := token.FindToken(payloadAsString)
		if tokenCandidate == nil {
			return "", fmt.Errorf("Security welcome did not find any token")
		}
		return *tokenCandidate, nil
	}

	return "", fmt.Errorf("Something screwed up, not redirected in welcome security")
}

func DoStrangeSecurityWelcome(txtParam string, cookies http.CookieJar) (string, error) {

	reqPayload := map[string]string{
		postFormFields.PdTxtParam: fmt.Sprintf(`txtParam=%s`, txtParam),
	}
	req, errCreatingReq := utils.CreateMultipartReq(reqPayload, eBankURLs.SecurityWelcome)

	if errCreatingReq != nil {
		return "", errCreatingReq
	}

	redirected := false
	httpClient := http.Client{
		Jar: cookies,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			redirected = true
			log.Printf("DoStrangeSecurityWelcome doing redirect to %s", req.URL.String())
			return nil //http.ErrUseLastResponse
		},
	}

	resp, errDoingReq := httpClient.Do(req)
	if errDoingReq != nil {
		return "", errDoingReq
	}
	payload, errReadingPayload := ioutil.ReadAll(resp.Body)
	if errReadingPayload != nil {
		return "", errReadingPayload
	}
	payloadAsString := string(payload)
	tokenCandidate := token.FindToken(payloadAsString)
	if tokenCandidate == nil {
		return "", fmt.Errorf("DoStrangeSecurityWelcome welcome did not find any token")
	}
	return *tokenCandidate, nil
	//if redirected {

	//}

	return "", fmt.Errorf("Something screwed up, not redirected in DoStrangeSecurityWelcome")
}

var postFormFields = utils.GetFieldNames()
var postFormValues = utils.GetFieldValues()
