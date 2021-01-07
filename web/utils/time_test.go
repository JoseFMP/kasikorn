package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type parseDateStringTestCase struct {
	asString      string
	expectedValue time.Time
}

func TestParseTimeCorrect(t *testing.T) {

	testCases := []*parseDateStringTestCase{
		&parseDateStringTestCase{
			asString:      "06/01/2021 02:07:47",
			expectedValue: time.Unix(1609873667, 0).In(GetThailandTimeZone()),
		},
	}

	for _, tc := range testCases {
		doTestCase(t, tc)
	}
}

func doTestCase(t *testing.T, tc *parseDateStringTestCase) {

	res, errParsing := PlainDateTimeStringToThaiTime(tc.asString)

	require.Nil(t, errParsing)
	require.Equal(t, tc.expectedValue, res)
}
