package statements

import "strings"

func findDownloadCommandInPayload(payload string) string {

	if strings.Contains(payload, postFormValues.ActionDownloadSa) {
		return postFormValues.ActionDownloadSa
	}
	return postFormValues.ActionDownloadCa
}

