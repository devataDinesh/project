package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"project/cache"
)

func GetList(r *http.Request) (map[string]bool, error) {
	userInput := cache.Input{}

	websiteList := cache.WebsiteList

	err := json.NewDecoder(r.Body).Decode(&userInput)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for _, val := range userInput.Website {
		websiteList[val] = false
	}
	return websiteList, nil
}
