package utils

import (
	"fmt"
	"io"
	"io/ioutil"

	"golang.org/x/net/html/charset"
)

func DecodeBody(body io.ReadCloser, contentType string) (string, error) {
	if contentType == "" {
		return "", fmt.Errorf("No content type specified")
	}
	reader, errCreatingReader := charset.NewReader(body, contentType)
	if errCreatingReader != nil {
		return "", errCreatingReader
	}
	asBytes, errReadingStream := ioutil.ReadAll(reader)
	if errReadingStream != nil {
		return "", errReadingStream
	}

	return string(asBytes), nil
}
