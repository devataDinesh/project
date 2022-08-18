package main

import (
	"fmt"
	"github.com/fatih/color"
	"net/http"
	"project/cache"
	"project/helper"
)

var list = cache.WebsiteList

func main() {

	color.Blue("Server Started Listening and Serving!!!...")

	http.HandleFunc("/GetList", getListHandler)

	http.HandleFunc("/CheckStatus", checkStatusHandler)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		return
	}
}

func getListHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	list, err = helper.GetList(r)

	if err != nil {
		color.Red("\t400 BAD REQUEST\n...........Failed to retrieve the WebsiteList :(...........")
	} else {
		fmt.Fprint(w, list)

		color.Green("\t200 OK\n...........Successfully retrieved the WebsiteList :)...........")
		color.Cyan("The received List is : ")

		for index, _ := range list {
			fmt.Println("\t", index)
		}
	}
}

func checkStatusHandler(w http.ResponseWriter, r *http.Request) {
}
