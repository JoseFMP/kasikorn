package statement

import (
	"testing"
	"time"

	"github.com/JoseFMP/kasikorn/web/utils"

	"github.com/JoseFMP/kasikorn/account"

	"github.com/stretchr/testify/require"
)

func TestCanParseMeta(t *testing.T) {

	mockRecords := [][]string{
		[]string{"Transaction History Detail As Of 02/12/2020"},
		[]string{"Account Number", "032-1-23971-4"},
		[]string{"Account Name", "N.O.O.N PROPERTY MANAGEMENT CO.,LTD."},
		[]string{"Type", "Savings Account"},
		[]string{"Statement Period", "01/12/2020 - 09/12/2020"},
	}

	meta, errParsing := parseRecordsMeta(mockRecords)

	require.Nil(t, errParsing)
	require.NotNil(t, meta)

	year, month, day := meta.AsOf.Date()
	require.Equal(t, 2020, year)
	require.Equal(t, time.Month(12), month)
	require.Equal(t, 2, day)

	require.NotNil(t, meta.Account)
	require.Equal(t, account.AccountNumber("032-1-23971-4"), meta.Account.Number)
	require.Equal(t, "N.O.O.N PROPERTY MANAGEMENT CO.,LTD.", meta.Account.Name)
	require.Equal(t, account.GetAllAccountTypes().Savings, meta.Account.Type)

	expectedFrom := time.Date(2020, time.December, 1, 0, 0, 0, 0, utils.GetThailandTimeZone())
	expectedFrom = expectedFrom.In(utils.GetThailandTimeZone())
	require.Equal(t, meta.Period.From.Unix(), expectedFrom.Unix())
	require.Equal(t, meta.Period.From, expectedFrom)

	expectedTo := time.Date(2020, time.December, 9, 0, 0, 0, 0, utils.GetThailandTimeZone())
	expectedTo = expectedTo.In(utils.GetThailandTimeZone())
	require.Equal(t, meta.Period.To.Unix(), expectedTo.Unix())
	require.Equal(t, meta.Period.To, expectedTo)

}

func TestCanParseBytes(t *testing.T) {

	bytesMock := []byte(mockStatement)
	records, errParsing := parseBytes(bytesMock)

	require.Nil(t, errParsing)
	require.NotNil(t, records)
	require.Len(t, records[0], 1)
	require.Len(t, records[1], 2)
	//require.Len(t, records[2], 2)
	//require.Len(t, records[3], 2)
	require.Len(t, records[4], 2)
	require.Len(t, records[7], 8)

}

const mockStatement = `
Transaction History Detail As Of 02/12/2020
Account Number,032-1-23971-4
Account Name,N.O.O.N PROPERTY MANAGEMENT CO.,LTD.   
Account Nickname,º¨¡. àÍç¹.âÍ.âÍ.àÍç¹ ¾ÃêÍ¾à¾ÍÃìµÕé àÁà¹¨àÁé¹
Type,Savings Account
Statement Period,01/10/2020 - 31/10/2020
Date,Transaction Type,Withdrawal (THB),Deposit (THB),Outstanding Balance (THB),Service Channel,Note
01/10/2020 09:04:40,"Cash Deposit NB","","10,000.00","389,319.99","HA YAEK CHALONG PHUKET BRANCH K0597253","",
01/10/2020 12:54:42,"Cheque/Money Transfer NB","","100,000.00","489,319.99","INTERNET BANKING AND MOBILE BANKING KCB08105","",
03/10/2020 15:22:09,"Transfer Withdrawal NB","40,000.00","","449,319.99","INTERNET BANKING AND MOBILE BANKING KCB08118","",
03/10/2020 16:06:23,"Cheque/Money Transfer NB","","40,000.00","489,319.99","PROMPTPAY FROM OTHER BANK IBA57065","",
08/10/2020 17:44:26,"Cheque/Money Transfer NB","","20,000.00","509,319.99","MOBILE BANKING KMP24128","",
10/10/2020 15:07:56,"Transfer Withdrawal NB","20,000.00","","489,319.99","INTERNET BANKING AND MOBILE BANKING KCB08107","",
17/10/2020 14:42:08,"Transfer Withdrawal NB","20,000.00","","469,319.99","INTERNET BANKING AND MOBILE BANKING KCB08112","",
19/10/2020 11:03:37,"Cash Deposit NB","","10,000.00","479,319.99","HA YAEK CHALONG PHUKET BRANCH K0719255","",
23/10/2020 11:13:31,"Cheque/Money Transfer NB","","6,000.00","485,319.99","PROMPTPAY FROM OTHER BANK IBA56014","",
24/10/2020 15:09:33,"Transfer Withdrawal NB","20,000.00","","465,319.99","INTERNET BANKING AND MOBILE BANKING KCB08116","",
26/10/2020 13:12:51,"Transfer Withdrawal NB","20,000.00","","445,319.99","INTERNET BANKING AND MOBILE BANKING KCB08119","",

`
