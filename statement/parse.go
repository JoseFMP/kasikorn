package statement

import (
	"bytes"
	"encoding/csv"
	"fmt"
)

const variableFieldsPerRecordValue = -1

func Parse(payload []byte) (*Statement, error) {
	records, errReadingRecords := parseBytes(payload)
	if errReadingRecords != nil {
		return nil, errReadingRecords
	}

	recordsMeta, errParsingMeta := parseRecordsMeta(records)
	if errParsingMeta != nil {
		return nil, errParsingMeta
	}
	return &Statement{
		Meta: *recordsMeta,
	}, nil
}

func parseBytes(payload []byte) ([][]string, error) {
	if len(payload) == 0 {
		return nil, fmt.Errorf("No payload provided ;)")
	}

	r := csv.NewReader(bytes.NewReader(payload))
	r.FieldsPerRecord = variableFieldsPerRecordValue // variable fields per record

	records, errReadingRecords := r.ReadAll()
	if errReadingRecords != nil {
		return nil, errReadingRecords
	}
	return records, nil
}

type MetaRowHeaders struct {
	TransactionHistory string
	AccountNumber      string
	AccountName        string
	AccountNickName    string
	Type               string
	StatementPeriod    string
}

func getMetaRowHeaders() MetaRowHeaders {
	return MetaRowHeaders{
		TransactionHistory: "Transaction History Detail As Of",
		AccountNumber:      "Account Number",
		AccountName:        "Account Name",
		AccountNickName:    "Account Nickname",
		Type:               "Type",
		StatementPeriod:    "Statement Period",
	}
}

var rowHeaders = getMetaRowHeaders()
