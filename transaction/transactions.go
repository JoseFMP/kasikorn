package transaction

import (
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
