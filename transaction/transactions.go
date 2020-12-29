package transaction

import (
	"strconv"
	"strings"
	"time"
)

type Transaction struct {
	Date                  time.Time       `json:"date"`
	Type                  TransactionType `json:"type"`
	AmountTHB             float64         `json:"amountTHB"`
	OutstandingBalanceTHB float64         `json:"outstandingBalanceTHB"`
	Channel               ServiceChannel  `json:"channel"`
	Note                  string          `json:"note"`
	ChequeNumber          *string         `json:"chequeNumber"`
}

type ServiceChannel string

func parseKasikornAmount(amountAsString string) (float64, error) {

	cleanedAmount := strings.ReplaceAll(amountAsString, ",", "")

	amount, errParsing := strconv.ParseFloat(cleanedAmount, 64)
	if errParsing != nil {
		return 0, errParsing
	}
	return amount, nil
}
