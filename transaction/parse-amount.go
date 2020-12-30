package transaction

import (
	"fmt"
	"strconv"
	"strings"
)

func parseWithdrawalOrDeposit(withDrawalAsString string, depositAsString string) (float64, error) {
	withdrawalCleaned := cleanAmount(withDrawalAsString)

	deposit := false
	amountAsString := withdrawalCleaned

	if withdrawalCleaned == "" {
		deposit = true
		depositCleaned := cleanAmount(depositAsString)
		if depositCleaned == "" {
			return 0, fmt.Errorf("No withdrawal and no deposit found, both empty")
		}
		amountAsString = depositAsString
	}

	var errParsing error
	amount, errParsing := parseKasikornAmount(amountAsString)
	if errParsing != nil {
		return 0, fmt.Errorf("Cannot parse amount. amountAsString:%s\nError:%v", amountAsString, errParsing)
	}
	if !deposit {
		amount = amount * (-1)
	}

	return amount, nil
}

func cleanAmount(amountToClean string) string {
	cleanedAmount := strings.ReplaceAll(amountToClean, "", "")
	cleanedAmount = strings.ReplaceAll(cleanedAmount, "\n", "")
	return cleanedAmount
}

func parseKasikornAmount(amountAsString string) (float64, error) {

	cleanedAmount := strings.ReplaceAll(amountAsString, ",", "")

	amount, errParsing := strconv.ParseFloat(cleanedAmount, 64)
	if errParsing != nil {
		return 0, errParsing
	}
	return amount, nil
}
