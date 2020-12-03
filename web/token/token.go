package token

import (
	"fmt"
	"regexp"
	"strings"
)

const Name = `org.apache.struts.taglib.html.TOKEN`

func getHTMLTokenPattern() string {
	return fmt.Sprintf(`<input type="hidden" name="%s" value="`, Name)
}

var htmlPattern = getHTMLTokenPattern()

func FindToken(loadToInspect string) *string {

	startDelimiter := strings.Index(loadToInspect, htmlPattern)
	if startDelimiter == -1 {
		return nil
	}
	token := loadToInspect[(startDelimiter + len(htmlPattern)):]
	endDelimiter := strings.Index(token, `"`)
	if endDelimiter == -1 {
		return nil
	}
	token = token[:endDelimiter]
	return &token
}

const patternTokenId = `<input type="hidden" name="tokenId" id="tokenId" value="`

var regexpTokenId = regexp.MustCompile(patternTokenId)

func FindTokenId(payloadToScan string) *string {

	targetBoundaries := regexpTokenId.FindStringIndex(payloadToScan)
	if targetBoundaries == nil {
		return nil
	}
	targetString := payloadToScan[(targetBoundaries[1] + 0):]
	endOfTheTokenDelimiter := strings.Index(targetString, `"/>`)
	targetString = targetString[:endOfTheTokenDelimiter]
	return &targetString
}
