package statements

import (
	"testing"

	"dev.azure.com/noon-homa/Kasikorn/_git/kasikorn/account"
	"github.com/stretchr/testify/assert"
)

func TestCanStrangeFormatAccountForDownload(t *testing.T) {

	mockAccountID := account.AccountID(`20161006245912`)
	mockAccountNumber := account.AccountNumber(`017-3-28346-6`)
	accountIdentificator := AccountIdentificator{
		Number: mockAccountNumber,
		ID:     mockAccountID,
	}
	formattedString := formatAccountNumberStrange(accountIdentificator)

	expectedResult := `20161006245912|0173283466|null|D|1152|0|null|THB|`

	assert.Equal(t, expectedResult, formattedString)

}
