package login

import (
	"fmt"
	"log"
	"net/http"

	"dev.azure.com/noon-homa/Kasikorn/kasikorn.git/web/token"
	"dev.azure.com/noon-homa/Kasikorn/kasikorn.git/web/utils"
)

func DoLogging(cookiesJar http.CookieJar, tokenId string, userName string, password string) error {
	payload := getLoggingDoPayload(userName, password, tokenId)
	req, errCreatingReq := utils.CreatePostFormReq(payload, kasikornURLs.Login)
	if errCreatingReq != nil {
		return errCreatingReq
	}
	setLoggingDoReqHeaders(&req.Header)
	redirected := false
	httpClient := http.Client{
		Jar: cookiesJar,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			redirected = true
			return nil
		},
	}
	resp, errMakingReq := httpClient.Do(req)
	if errMakingReq != nil {
		return errMakingReq
	}
	if redirected {

		//payload, errReadingData := ioutil.ReadAll(resp.Body)
		//if errReadingData != nil {
		//	return errReadingData
		//}
		payloadAsString, errDecodingPayload := utils.DecodeBody(resp.Body, resp.Header.Get("Content-Type"))
		if errDecodingPayload != nil {
			return errDecodingPayload
		}
		candidateToken := token.FindToken(payloadAsString)
		if candidateToken != nil {
			log.Println(candidateToken)
		}
		return nil
	}
	return fmt.Errorf("Request was not redirected. Strange!")
}

func getLoggingDoPayload(userName string, password string, tokenId string) map[string]string {
	payload := map[string]string{
		"locale":   "en",
		"app":      "0",
		"userName": userName,
		"password": password,
		"custType": utils.SmeCustomerType,
		"tokenId":  tokenId,
		"cmd":      "authenticate",
	}
	return payload
}

func setLoggingDoReqHeaders(header *http.Header) {
	header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	header.Add("Accept-Encoding", "gzip, deflate, br")
	header.Add("Accept-Language", "en-US,en;q=0.9,es;q=0.8,fr;q=0.7,th;q=0.6,de;q=0.5")
	header.Add("Cache-Control", "max-age=0")
	header.Add("Connection", "close")
	header.Add("Host", kasikornURLs.HostName)
	header.Add("Origin", kasikornURLs.Origin)
}

var eBankURLs = utils.GetAllEbankURLs()
var kasikornURLs = utils.GetAllOnlineKasikornURLs()
