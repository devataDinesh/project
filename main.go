package main

import (
	"github.com/fatih/color"
	"net/http"
)

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

}

func checkStatusHandler(w http.ResponseWriter, r *http.Request) {
}