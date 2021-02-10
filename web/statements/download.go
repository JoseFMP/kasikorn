package statements

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/JoseFMP/kasikorn/web/token"
	"github.com/JoseFMP/kasikorn/web/utils"
)

func RequestDownload(from utils.KasikornDate,
	to utils.KasikornDate,
	account AccountIdentificator,
	downloadCommand string,
	cookies http.CookieJar, tokenToUse string) ([]byte, error) {

	requestPayload := getDownloadRequestPayload(from, to, account, tokenToUse, downloadCommand)
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

func getDownloadRequestPayload(from utils.KasikornDate, to utils.KasikornDate, account AccountIdentificator, tokenToSend string, downloadCommand string) map[string]string {
	values := map[string]string{
		token.Name:                   tokenToSend,
		postFormFields.AccountNo:     string(account.ID),
		postFormFields.AccountNumber: string(account.Number),
		postFormFields.SelAccountNo:  formatAccountNumberStrange(account),
		postFormFields.AccountDesc:   "",
		postFormFields.Action:        downloadCommand,
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
func formatAccountNumberStrange(account AccountIdentificator) string {

	accountNumberNoDashes := strings.ReplaceAll(string(account.Number), "-", "")
	return fmt.Sprintf(`%s|%s|null|D|1152|0|null|THB|`, account.ID, accountNumberNoDashes)
}
