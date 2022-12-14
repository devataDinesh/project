package main

//Problem Statement:

//We are going to build an application that is meant to be a sort of status checker for some
//common websites that exist online. The application continuously polls the status of these websites
//and exposes APIs to retrieve the information.

import (
	"fmt"
	"github.com/fatih/color"
	"net/http"
	"project/cache"
	"project/checker"
	"project/helper"
	"time"
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
	r.ParseForm()
	optionalQuery := r.Form.Get("name")
	if optionalQuery == "" {
		for {
			color.Cyan(fmt.Sprintf("\n\n" + time.Now().String() + "\n\n"))
			helper.GetStatusHelper(list)
			time.Sleep(15 * time.Second)
		}
	} else {
		color.Magenta(fmt.Sprintf("\nObtained Optional Query parameter is " + optionalQuery))
		_, ok := list[optionalQuery]
		if ok {
			checkerObj := checker.StatusChecker{}
			res, _ := checkerObj.CheckStatus(optionalQuery)
			if res {
				color.Green(fmt.Sprintf("The Website " + optionalQuery + " is UP"))
			} else {
				color.Red(fmt.Sprintf("The Website " + optionalQuery + " is DOWN"))
			}
		} else {
			color.Yellow("Sorry!... No such Website Provided... :(")
		}
	}
}
