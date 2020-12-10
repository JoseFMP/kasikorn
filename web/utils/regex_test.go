package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRegexpDateTime(t *testing.T) {

	exampleDateString := `02/12/2020 11:05:15 traaaalalalala`
	regexpToUse := GetRegexpSlashedDateAndTime()

	stringResult := regexpToUse.FindString(exampleDateString)
	require.NotEmpty(t, stringResult)

}

func TestRegexpDate(t *testing.T) {

	exampleDateString := `02/12/2020 11:05:15 traaaalalalala`
	regexpToUse := GetRegexpSlashedDate()

	stringResult := regexpToUse.FindString(exampleDateString)
	require.NotEmpty(t, stringResult)

}
