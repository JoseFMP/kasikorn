package transaction

import (
	"testing"
	"time"

	"dev.azure.com/noon-homa/Kasikorn/kasikorn.git/web/utils"

	"github.com/stretchr/testify/require"
)

func TestCanParseNormalRecords(t *testing.T) {

	record1 := []string{
		"02/12/2020 11:05:15", "Cheque/Money Transfer NB", "", "18,798.08", "193,128.71", "DOMESTIC MONEY TRANSFER BNT00001", "",
	}

	tr, errParsingTr := parseTransaction(record1, false)

	require.Nil(t, errParsingTr)
	require.NotNil(t, tr)

	expectedTimestamp := time.Date(2020, time.December, 2, 11, 5, 15, 0, utils.GetThailandTimeZone())
	expectedTimestamp = expectedTimestamp.In(utils.GetThailandTimeZone())

	require.Equal(t, expectedTimestamp.Unix(), tr.Date.Unix())
	require.Equal(t, expectedTimestamp, tr.Date)

	require.Equal(t, TransactionType(GetAllTransactionTypes().ChequeMoneyTransfer), tr.Type)
	require.Equal(t, ServiceChannel("DOMESTIC MONEY TRANSFER BNT00001"), tr.Channel)

	require.Equal(t, float64(18798.08), tr.AmountTHB)

}

func TestCanParseNoNBParticle(t *testing.T) {

	record1 := []string{
		"02/12/2020 11:05:15", "Cheque/Money Transfer", "", "18,798.08", "193,128.71", "DOMESTIC MONEY TRANSFER BNT00001", "",
	}

	tr, errParsingTr := parseTransaction(record1, false)

	require.Nil(t, errParsingTr)
	require.NotNil(t, tr)

	expectedTimestamp := time.Date(2020, time.December, 2, 11, 5, 15, 0, utils.GetThailandTimeZone())
	expectedTimestamp = expectedTimestamp.In(utils.GetThailandTimeZone())

	require.Equal(t, expectedTimestamp.Unix(), tr.Date.Unix())
	require.Equal(t, expectedTimestamp, tr.Date)

	require.Equal(t, TransactionType(GetAllTransactionTypes().ChequeMoneyTransfer), tr.Type)
	require.Equal(t, ServiceChannel("DOMESTIC MONEY TRANSFER BNT00001"), tr.Channel)

	require.Equal(t, float64(18798.08), tr.AmountTHB)

}
