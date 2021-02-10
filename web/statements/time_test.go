package statements

import (
	"testing"

	"github.com/JoseFMP/kasikorn/web/utils"
	"github.com/stretchr/testify/require"
)

func TestCanCreatePayload(t *testing.T) {

	//arrange
	payloadOutput := map[string]string{}

	from := utils.KasikornDate{
		Year: 2021,
		Day:  1,
	}
	to := utils.KasikornDate{
		Year: 2021,
		Day:  12,
	}

	// act
	setFromToIntoPayload(from, to, payloadOutput, true)

	//verify
	require.NotNil(t, payloadOutput)

	fromDay, fromMonth, fromYear := getFrom(payloadOutput)
	toDay, toMonth, toYear := getTo(payloadOutput)

	require.Equal(t, "2021", fromYear)
	require.Equal(t, "2021", toYear)

	require.Equal(t, "01", fromMonth)
	require.Equal(t, "01", toMonth)

	require.Equal(t, "01", fromDay)
	require.Equal(t, "12", toDay)
}

func getFrom(payload map[string]string) (string, string, string) {

	fromDay := payload[postFormFields.SelDayFrom]
	fromMonth := payload[postFormFields.SelMonthFrom]
	fromYear := payload[postFormFields.SelYearFrom]

	return fromDay, fromMonth, fromYear

}

func getTo(payload map[string]string) (string, string, string) {

	day := payload[postFormFields.SelDayTo]
	month := payload[postFormFields.SelMonthTo]
	year := payload[postFormFields.SelYearTo]

	return day, month, year
}
