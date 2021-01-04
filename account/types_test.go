package account

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type parseAccountTypeTestCase struct {
	accountTypeAsString string
	expectedValue       AccountType
}

func TestCanParseAccountType(t *testing.T) {
	testCases := []*parseAccountTypeTestCase{
		&parseAccountTypeTestCase{
			accountTypeAsString: "Savings account",
			expectedValue:       GetAllAccountTypes().Savings,
		},
	}
	for _, tc := range testCases {
		testParseTestCase(t, tc)
	}
}

func testParseTestCase(t *testing.T, testCase *parseAccountTypeTestCase) {

	result := GetAccountType(testCase.accountTypeAsString)

	//verify
	require.Equal(t, testCase.expectedValue, result)

}
