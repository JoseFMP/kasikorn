package statements

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"dev.azure.com/noon-homa/Kasikorn/_git/kasikorn/account"
	"dev.azure.com/noon-homa/Kasikorn/_git/kasikorn/web/token"
	"dev.azure.com/noon-homa/Kasikorn/_git/kasikorn/web/utils"
)

func SelectAccountForStatementInquiry(from time.Time, to time.Time, accountID account.AccountID, tokenToSend string, cookies http.CookieJar) (string, string, error) {
	payload := getSelectStatementPayload(from, to, accountID, tokenToSend)
	req, errCreatingReq := utils.CreatePostFormReq(payload, eBankURLs.StatementInquiry)
	if errCreatingReq != nil {
		return "", "", errCreatingReq
	}
	setHeadersSelectAccount(&req.Header)
	httpClient := http.Client{
		Jar: cookies,
	}

	resp, errDoingReq := httpClient.Do(req)
	if errDoingReq != nil {
		return "", "", errDoingReq
	}

	if resp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("Making first HTTP req returned not OK: %d - %s", resp.StatusCode, resp.Status)
	}

	responsePayload, errReadingResp := ioutil.ReadAll(resp.Body)
	if errReadingResp != nil {
		return "", "", errReadingResp
	}

	respAsString := string(responsePayload)
	errVerifyingPayload := verifyResponsePayloadSelect(respAsString)
	if errVerifyingPayload != nil {
		return "", "", errVerifyingPayload
	}

	tokenCandidate := token.FindToken(respAsString)
	if tokenCandidate != nil {
		downloadCommand := findDownloadCommandInPayload(respAsString)
		return *tokenCandidate, downloadCommand, nil
	} else {
		return "", "", fmt.Errorf("Did not find a token in selectAccountForStatementInquiry")
	}
}

func setHeadersSelectAccount(header *http.Header) {
	header.Add("Connection", "close")
	header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	header.Add("Cache-Control", "max-age=0")
	//header.Add("Accept-Encoding", "gzip, deflate, br")
	header.Add("Accept-Language", "en-US,en;q=0.9,es;q=0.8,fr;q=0.7,th;q=0.6,de;q=0.5")
	header.Add("Host", eBankURLs.Hostname)
	header.Add("Origin", eBankURLs.Origin)
	header.Add("Refer", eBankURLs.StatementInquiry)
	header.Add("Sec-Fetch-Dest", "iframe")
	header.Add("Sec-Fetch-Mode", "navigate")
	header.Add("Sec-Fetch-Site", "same-origin")
}

func getSelectStatementPayload(from time.Time, to time.Time, accountID account.AccountID, tokenToSend string) map[string]string {
	values := map[string]string{
		postFormFields.AccountDesc: "",
		postFormFields.Action:      postFormValues.ActionSelect,
		postFormFields.AccountNo:   string(accountID),
		postFormFields.Period:      "3",
		token.Name:                 tokenToSend,
	}

	setFromToIntoPayload(from, to, values, true)
	return values
}

func setFromToIntoPayload(from time.Time, to time.Time, payload map[string]string, includeDateFromDateTo bool) {
	fromDay, fromMonth, fromYear := timeToPayload(from)

	payload["selDayFrom"] = fromDay
	payload["selMonthFrom"] = fromMonth
	payload["selYearFrom"] = fromYear

	toDay, toMonth, toYear := timeToPayload(to)
	payload["selDayTo"] = toDay
	payload["selMonthTo"] = toMonth
	payload["selYearTo"] = toYear
	if includeDateFromDateTo {
		payload["selDateFrom"] = from.Format(selDateLayout)
		payload["selDateTo"] = to.Format(selDateLayout)
	}
}

const selDateLayout = "02/01/2006"

func timeToPayload(timeToEmbed time.Time) (string, string, string) {

	day := fmt.Sprintf("%02d", timeToEmbed.Day())
	month := fmt.Sprintf("%02d", timeToEmbed.Month())
	year := fmt.Sprintf("%02d", timeToEmbed.Year())

	return day, month, year
}

const inquiryResultPattern = `<h1>Statement Inquiry - Inquiry Result</h1>`

func verifyResponsePayloadSelect(payload string) error {
	errorResponse := verifyErrorResponse(payload)
	if errorResponse != nil {
		return errorResponse
	}
	containsInquiryResult := strings.Contains(payload, inquiryResultPattern)
	if !containsInquiryResult {
		return fmt.Errorf("Does not contain inquiryResult")
	}
	return nil
}

func getErrorStrings() []string {
	return []string{
		"Sorry, the current logged-in session",
		"Sorry, your request cannot be processed",
		"An unexpected response has been detected",
		"Select for last 6 month and before today only",
		"only ONCE",
	}
}

func verifyErrorResponse(payload string) error {
	if payload == "" {
		return fmt.Errorf("Payload is empty")
	}

	for _, es := range getErrorStrings() {
		if strings.Contains(payload, es) {
			return fmt.Errorf(es)
		}
	}

	return nil
}

var postFormFields = utils.GetFieldNames()
var postFormValues = utils.GetFieldValues()

type AccountIdentificator struct {
	Number account.AccountNumber
	ID     account.AccountID
}
