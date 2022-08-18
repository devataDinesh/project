package cache

type Input struct {
	Website []string `json:"website"`
}

var WebsiteList = make(map[string]bool)
