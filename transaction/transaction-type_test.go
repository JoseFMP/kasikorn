package transaction

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseWithholdingType(t *testing.T) {

	testCase := transactionTypeTestCase{
		asString:      "Withholding Tax Payable NB",
		expectedValue: GetAllTransactionTypes().WitholdingTaxPayable,
	}

	canParseTransactionType(t, testCase)

}

type transactionTypeTestCase struct {
	asString      string
	expectedValue TransactionType
}

func canParseTransactionType(t *testing.T, testCase transactionTypeTestCase) {

	result := transactionTypeFromString(testCase.asString)

	require.Equal(t, testCase.expectedValue, *result)
}
