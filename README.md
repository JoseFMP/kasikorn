[![Build status](https://dev.azure.com/noon-homa/Kasikorn/_apis/build/status/kasikorn-ci)](https://dev.azure.com/noon-homa/Kasikorn/_build/latest?definitionId=27)

# Motivation
Traditional banks are well known to be stale in terms of IT infrastructure and services. Very often, they do not expose any public (or private) APIs. 

Kasikorn bank is not different. It does not offer any proper way of programatically retrieve statements or balance of banks accounts it is holding. I therefore developed this Go-based client library to leverage the web portal of Kasikorn to download and parse statements from Kasikorn.

Virtually every business has to check the statements of its bank accounts regularily. Common tasks include to acknowledge client payments, check accounts balance, or prepare reports of the accounts transactions for accountans (i.e. tax reports). The number of people who need to access this information is usually high (directors, accountans, HR, sales).


# How does it work?
I used an interceptor of HTTP requests / responses / redirects / cookies while navigating through the SME web portal of Kasikorn bank. Based on the captures I reverse-engineered the logic of the HTTP client and put it in a Go library. By using the function `Login()` you will be indeed navigating the web portal of Kasikorn entering your credentials. You can then download statements based on dates.


# Contributing
This is a library written in Go to interact with Kasikorn bank. This project is not currently very actively developed as I only implemented some chunks that I needed to be able to authenticate with Kasikorn through their web interface and download statements of accounts.

If you are interested in colaborate get in touch at jose@mingorance.engineer

# Get started

Plug (in a secure way) your login credentials into a `config.Config` struct:
```
configuration := config.Config{
		UserName: username,
		Password: password,
	}
```
WARN: Do not hardcode or share your creds! That said, as it is now with this credentials you have only read-access, i.e. MFA is fortunately in.

Just init a session:
```
	session := web.InitSession(configuration)
```

What this does is to navigate through the Kasikorn bank web portal and try to login using your credentials. And indeed verifying the credentials actually work and a statement can be downloaded.

Download a statement: 
```
statement, errGettingStatement := session.GetStatement(from, to, accountNumber)
```

Note that `from` and `to` are not Go `time.Time` objects but custom struct in which two coordinates are defined, `Year` and `Day`. Indeed, `from` and `to` are not timestamps but "calendar days" based on calendar days defined by Kasikorn, i.e. time zone is GMT + 7. Currently there is no granularity to request through the portal statements by time, i.e. smallest time window is 1 day.
 

