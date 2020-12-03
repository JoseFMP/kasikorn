package utils

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"net/url"
	"strings"
)

const contenTypeHttpHeaderName = `Content-Type`
const contentTypeURLEncoded = `application/x-www-form-urlencoded`

func CreatePostFormReq(payload map[string]string, targetURL string) (*http.Request, error) {
	formDataAsString := encodePayload(payload)
	req, errCreatingReq := http.NewRequest(http.MethodPost, targetURL, strings.NewReader(formDataAsString))

	if errCreatingReq != nil {
		return nil, errCreatingReq
	}

	req.Header.Add(contenTypeHttpHeaderName, contentTypeURLEncoded)

	return req, nil
}

func CreateMultipartReq(payload map[string]string, targetURL string) (*http.Request, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, _ := writer.CreatePart(textproto.MIMEHeader{"Content-Type": {"application/json"}})
	formDataAsString := encodePayload(payload)

	_, err := part.Write([]byte(formDataAsString))
	if err != nil {
		return nil, err
	}
	writer.Close()

	req, _ := http.NewRequest(http.MethodPost, targetURL, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, nil
}

func encodePayload(payload map[string]string) string {
	formData := url.Values{}

	for k, v := range payload {
		formData.Set(k, v)
	}
	return formData.Encode()
}

type postFieldValues struct {
	ActionDownloadSa string
	ActionDownloadCa string
	ActionSelect     string
}

type postFieldNames struct {
	AccountNo     string
	AccountNumber string
	AccountDesc   string
	No            string
	SelAccountNo  string
	Action        string
	Period        string
	SelPrevMonth  string
	PdTxtParam    string
	TxtParam      string
}

func GetFieldNames() postFieldNames {
	return postFieldNames{
		AccountDesc:   "accountDesc",
		No:            "no",
		AccountNo:     "accountNo",
		AccountNumber: "accountNumber",
		SelAccountNo:  "selAccountNo",
		Action:        "action",
		Period:        "period",
		SelPrevMonth:  "selPrevMonth",
		PdTxtParam:    "_pd",
		TxtParam:      "txtParam",
	}
}

func GetFieldValues() postFieldValues {
	return postFieldValues{
		ActionDownloadSa: "sa_download",
		ActionDownloadCa: "ca_download",
		ActionSelect:     "select",
	}
}
