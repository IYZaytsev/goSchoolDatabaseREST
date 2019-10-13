package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//StudentCreation takes in a form value from the other HTTP server and creates a new student
//Which is then sends back in response
func StudentCreation(w http.ResponseWriter, r *http.Request) {

	body, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	parameter := string(body)
	studentStruct := School.AddStudent(parameter)
	w.Header().Set("Content-type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(studentStruct); err != nil {
		panic(err)
	}

}

//TeacherCreation takes in a form value from the other HTTP server and creates a new Teacher
//Which is then sends back in response
func TeacherCreation(w http.ResponseWriter, r *http.Request) {

	body, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	parameter := string(body)

	teacherStruct := School.AddTeacher(parameter)

	w.Header().Set("Content-type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(teacherStruct); err != nil {
		panic(err)
	}

}

//ClassCreation takes in a form value from the other HTTP server and creates a new Class
//Which is then sends back in response
func ClassCreation(w http.ResponseWriter, r *http.Request) {

	body, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	parameter := string(body)

	classStruct := School.AddClass(parameter)

	w.Header().Set("Content-type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(classStruct); err != nil {
		panic(err)
	}

}
