package statements

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"dev.azure.com/noon-homa/Kasikorn/_git/kasikorn/account"
	"dev.azure.com/noon-homa/Kasikorn/_git/kasikorn/web/token"
	"dev.azure.com/noon-homa/Kasikorn/_git/kasikorn/web/utils"
)

func RequestDownload(from time.Time,
	to time.Time,
	account account.Account, cookies http.CookieJar, tokenToUse string) ([]byte, error) {

	requestPayload := getDownloadRequestPayload(from, to, account.ID, account.Number, tokenToUse)
	req, errCreatingReq := utils.CreatePostFormReq(requestPayload, eBankURLs.StatementInquiry)
	if errCreatingReq != nil {
		return nil, errCreatingReq
	}

	setDownloadRequestHeaders(&req.Header)

	httpClient := http.Client{
		Jar: cookies,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			log.Println("Redirected!")
			return nil
		},
	}

	resp, errDoingReq := httpClient.Do(req)

	if errDoingReq != nil {
		return nil, errDoingReq
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP request to download statement received response not ok: %d - %s", resp.StatusCode, resp.Status)
	}

	contentDispositionHeader := resp.Header.Get("Content-Disposition")
	if contentDispositionHeader != "attachment; filename=download.csv" {
		return nil, fmt.Errorf("Download response does not have a right Content-Disposition header: %s", contentDispositionHeader)
	}

	responsePayload, errReadingResp := ioutil.ReadAll(resp.Body)
	if errReadingResp != nil {
		return nil, errReadingResp
	}

	return responsePayload, nil
}

var eBankURLs = utils.GetAllEbankURLs()

func getDownloadRequestPayload(from time.Time, to time.Time, accountID string, accountNumber string, tokenToSend string) map[string]string {
	values := map[string]string{
		token.Name:                   tokenToSend,
		postFormFields.AccountNo:     accountID,
		postFormFields.AccountNumber: accountNumber,
		postFormFields.SelAccountNo:  formatAccountNumberStrange(accountID, accountNumber),
		postFormFields.AccountDesc:   "",
		postFormFields.Action:        postFormValues.ActionDownloadCa,
		postFormFields.SelPrevMonth:  "",
		postFormFields.Period:        "3",
		postFormFields.No:            "0",
	}
	setFromToIntoPayload(from, to, values, false)

	return values
}

func setDownloadRequestHeaders(header *http.Header) {
	//header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	//header.Set("Accept-Encoding", "gzip, deflate, br")
	header.Set("Accept-Language", "en-US,en;q=0.9,es;q=0.8,fr;q=0.7,th;q=0.6,de;q=0.5")

	header.Set("Host", eBankURLs.Hostname)
	header.Set("Origin", eBankURLs.Origin)
	header.Set("Cache-Control", "max-age=0")
	header.Set("Connection", "close")
	header.Set("Refer", eBankURLs.StatementInquiry)
	header.Set("Sec-Fetch-Dest", "iframe")
	header.Set("Sec-Fetch-Mode", "navigate")
	header.Set("Sec-Fetch-Site", "same-origin")
	header.Set("Sec-Fetch-User", "?1")
}

//20161006245912|0173283466|null|D|1152|0|null|THB|
func formatAccountNumberStrange(accountID string, accountNumber string) string {

	accountNumberNoDashes := strings.ReplaceAll(accountNumber, "-", "")
	return fmt.Sprintf(`%s|%s|null|D|1152|0|null|THB|`, accountID, accountNumberNoDashes)

	//return fmt.Sprintf("|%s||||||", accountNumberNoDashes)
}
