package statements

import (
	"testing"
	"time"

	"dev.azure.com/noon-homa/Kasikorn/kasikorn.git/account"
	"github.com/stretchr/testify/assert"
)

func TestCanStrangeFormatAccountForDownload(t *testing.T) {

	mockAccountID := account.AccountID(`20161006245912`)
	mockAccountNumber := account.AccountNumber(`017-3-28346-6`)
	formattedString := formatAccountNumberStrange(mockAccountID, mockAccountNumber)

	expectedResult := `20161006245912|0173283466|null|D|1152|0|null|THB|`

	assert.Equal(t, expectedResult, formattedString)

}

func TestCanComposePayloadDays(t *testing.T) {

	mockTime, _ := time.Parse(selDateLayout, "06/01/1987")

	day, month, year := timeToPayload(mockTime)

	assert.Equal(t, "06", day)
	assert.Equal(t, "01", month)
	assert.Equal(t, "1987", year)

}
