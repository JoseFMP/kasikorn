[![Build status](https://dev.azure.com/noon-homa/Kasikorn/_apis/build/status/kasikorn-ci)](https://dev.azure.com/noon-homa/Kasikorn/_build/latest?definitionId=27)

# Motivation
Virtually every business of at least an average size, has to check the statements of its bank account regularily. Common tasks include to verify if a client payment has really arrived, what the current balance is, or the more tedious one, provide information for the accountants so that they can move forward. The number of people required to access the statements is significant and the frequency usually high.

Traditional banks are well known to be stale in terms of IT infrastructure and services. Very often, they do not expose any public (or private) API. 

Kasikorn is not different. It does not offer any proper way of programatically figure out the statements or balance of your accounts. I therefore developed this Go-based client library to leverage the web portal of Kasikorn to download and parse statements from Kasikorn.

# How does it work?
I used an interceptor of HTTP requests / responses / redirects / cookies while navigating through the SME web portal of Kasikorn. Based on the captures I reversed engineer the logic of the web client and put it in a Go library. By using the function `Login()` you will be indeed navigating the web portal of Kasikorn entering your credentials. You can then download statements based on dates.


# Contributing
This is a library written in Go to interact with Kasikorn bank. This project is not currently very actively developed as I only implemented some chunks that I needed to be able to authenticate with Kasikorn through their web interface and download statements of accounts.

If you are interested in colaborate get in touch at jose@mingorance.engineer
