package login

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"dev.azure.com/noon-homa/Kasikorn/kasikorn.git/web/token"
	"dev.azure.com/noon-homa/Kasikorn/kasikorn.git/web/utils"
)

func DoPreloggin(cookiesJar http.CookieJar) (string, error) {
	req, errCreatingReq := http.NewRequest(http.MethodGet, kasikornURLs.Prelogging, nil)
	if errCreatingReq != nil {
		return "", errCreatingReq
	}
	getParams := url.Values{}
	getParams.Set("lang", "EN")
	getParams.Set("type", utils.SmeCustomerType)

	req.URL.RawQuery = getParams.Encode()

	httpClient := http.Client{
		Jar: cookiesJar,
	}
	response, errMakingReq := httpClient.Do(req)
	if errMakingReq != nil {
		return "", errMakingReq
	}
	respPayload, errReadingBody := ioutil.ReadAll(response.Body)
	if errReadingBody != nil {
		return "", errReadingBody
	}
	payloadAsString := string(respPayload)
	tokenId := token.FindTokenId(payloadAsString)
	if tokenId == nil {
		return "", fmt.Errorf("Could not find the token ID!")
	}
	log.Printf("Found token ID: %s", *tokenId)
	return *tokenId, nil
}
