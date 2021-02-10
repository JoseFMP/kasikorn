package statement

import (
	"bytes"
	"encoding/csv"
	"fmt"

	"github.com/JoseFMP/kasikorn/transaction"
)

const variableFieldsPerRecordValue = -1

// Parse receives an argument raw bytes of a statement file from Kasikorn, i.e. CSV file from Kasikorn and parses it.
func Parse(payload []byte) (*Statement, error) {
	records, errReadingRecords := parseBytes(payload)
	if errReadingRecords != nil {
		return nil, errReadingRecords
	}

	recordsMeta, errParsingMeta := parseRecordsMeta(records)
	if errParsingMeta != nil {
		return nil, errParsingMeta
	}

	transactionFields := records[6]
	hasCheckNumber := true
	if len(transactionFields) == 7 {
		hasCheckNumber = false
	}
	transactions := transaction.Parse(records[7:], hasCheckNumber)
	return &Statement{
		Meta:         *recordsMeta,
		Transactions: transactions,
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
