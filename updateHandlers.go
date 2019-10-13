package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func StudentUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	//Reads body from the request seperates into 2 different variable ID and Class ID
	body, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	parameter := string(body)

}
