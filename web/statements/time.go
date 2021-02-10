package statements

import (
	"fmt"

	"github.com/JoseFMP/kasikorn/web/utils"
)

func setFromToIntoPayload(from utils.KasikornDate, to utils.KasikornDate, payload map[string]string, includeDateFromDateTo bool) {

	fromDay, fromMonth, fromYear := from.ToChunks()

	payload[postFormFields.SelDayFrom] = fromDay
	payload[postFormFields.SelMonthFrom] = fromMonth
	payload[postFormFields.SelYearFrom] = fromYear

	toDay, toMonth, toYear := to.ToChunks()
	payload[postFormFields.SelDayTo] = toDay
	payload[postFormFields.SelMonthTo] = toMonth
	payload[postFormFields.SelYearTo] = toYear
	if includeDateFromDateTo {
		payload[postFormFields.SelDateFrom] = formatChunks(fromDay, fromMonth, fromYear)
		payload[postFormFields.SelDateTo] = formatChunks(toDay, toMonth, toYear)
	}
}

func formatChunks(day string, month string, year string) string {
	return fmt.Sprintf("%s/%s/%s", day, month, year)
}
