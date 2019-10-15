package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

//StudentUpdate updates students class and anme
func StudentUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	//Reads body from the request seperates into 2 different variable ID and Class ID
	body, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	parameter := string(body)
	searchCriteria := strings.Split(parameter, "/")
	className, studentName := searchCriteria[0], searchCriteria[2]
	studentID, _ := strconv.Atoi(searchCriteria[1])
	School.UpdateStudent(studentID, studentName, className)

}

//TeacherUpdate updates teachers class and name
func TeacherUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	//Reads body from the request seperates into 2 different variable ID and Class ID
	body, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	parameter := string(body)

	searchCriteria := strings.Split(parameter, "/")
	className, teacherName := searchCriteria[0], searchCriteria[2]
	teacherID, _ := strconv.Atoi(searchCriteria[1])
	School.UpdateTeacher(teacherID, teacherName, className)

}

//ClassUpdate updates className
func ClassUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	//Reads body from the request seperates into 2 different variable ID and Class ID
	body, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	parameter := string(body)

	searchCriteria := strings.Split(parameter, "/")
	classID, _ := strconv.Atoi(searchCriteria[0])
	className := searchCriteria[1]

	School.UpdateClass(classID, className)

}
