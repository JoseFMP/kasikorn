package utils

import (
	"fmt"
)

type onlineKasikornURLs struct {
	HostName      string
	Origin        string
	Base          string
	CheckSession  string
	RetailWelcome string
	Home          string
	Login         string
	RedirectIB    string
	Prelogging    string
}

type ebankKasikornURLs struct {
	Hostname         string
	Origin           string
	SecurityWelcome  string
	Base             string
	StatementInquiry string
}

func GetAllEbankURLs() ebankKasikornURLs {

	hostname := `ebank.kasikornbankgroup.com`
	origin := fmt.Sprintf("https://%s", hostname)
	base := fmt.Sprintf("%s/retail", origin)

	securityWelcomeSubpath := `security/Welcome.do`
	statementInquirySubpath := "accountinfo/AccountStatementInquiry.do"

	return ebankKasikornURLs{
		Hostname:         hostname,
		Origin:           origin,
		Base:             base,
		SecurityWelcome:  fmt.Sprintf("%s/%s", base, securityWelcomeSubpath),
		StatementInquiry: fmt.Sprintf("%s/%s", base, statementInquirySubpath),
	}
}

func GetAllOnlineKasikornURLs() onlineKasikornURLs {

	hostname := `online.kasikornbankgroup.com`
	origin := fmt.Sprintf("https://%s", hostname)
	baseURL := fmt.Sprintf("%s/K-Online", origin)

	retailWelcomeSubpath := `retail/RetailWelcome.do`
	homeSubpath := `indexHome.jsp`
	checkSessionSubpath := `checkSession.jsp`
	loginSubpath := `login.do`
	redirectIBSubpath := "ib/redirectToIB.jsp"
	preloggingSubpath := `login.jsp`

	return onlineKasikornURLs{
		HostName:      hostname,
		Origin:        origin,
		Base:          baseURL,
		CheckSession:  fmt.Sprintf("%s/%s", baseURL, checkSessionSubpath),
		RetailWelcome: fmt.Sprintf("%s/%s", baseURL, retailWelcomeSubpath),
		Home:          fmt.Sprintf("%s/%s", baseURL, homeSubpath),
		Login:         fmt.Sprintf("%s/%s", baseURL, loginSubpath),
		RedirectIB:    fmt.Sprintf("%s/%s", baseURL, redirectIBSubpath),
		Prelogging:    fmt.Sprintf("%s/%s", baseURL, preloggingSubpath),
	}
}

const SmeCustomerType = "sme"
