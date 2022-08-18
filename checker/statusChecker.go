package checker

import (
	"net/http"
)

type StatusChecker struct{}

func (s *StatusChecker) CheckStatus(websiteName string) (bool, error) {
	response, err := http.Get("https://" + websiteName + "/")
	if err != nil {
		return false, err
	}
	return response.StatusCode == 200, nil
}
