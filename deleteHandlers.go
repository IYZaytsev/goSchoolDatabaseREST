package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func StudentDelete(w http.ResponseWriter, r *http.Request) {

	//gets passed ID for deletion
	body, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	parameter := string(body)
	fmt.Println(parameter)
	studentID, _ := strconv.Atoi(parameter)
	School.DeleteStudent(studentID)
}
