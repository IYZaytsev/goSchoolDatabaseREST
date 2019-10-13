package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	schoolLib "./schoolLibrary"
)

var myClient = http.Client{Timeout: 10 * time.Second}
var dataBaseStudent schoolLib.Student

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func ClassIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	classesId := School.GetClasses()

	if err := json.NewEncoder(w).Encode(classesId); err != nil {
		panic(err)
	}
}

//StudentSearch sends a slice of all students who match a search criteria
//if search criteria is empty returns all students
func StudentSearch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	//Reads body from the request seperates into 3 different variable
	body, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	parameter := string(body)
	searchCriteria := strings.Split(parameter, "/")
	name, class, id := searchCriteria[0], searchCriteria[1], searchCriteria[2]

	//if all search criteria are empty then it sends a slice of all students
	if len(name) == 0 && len(class) == 0 && len(id) == 0 {

		studentSlice := School.GetStudents()
		w.Header().Set("Content-type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(studentSlice); err != nil {
			panic(err)
		}

	} else {
		//else it sends search results
		studentSlice := School.SearchStudents(name, class, id)

		if err := json.NewEncoder(w).Encode(studentSlice); err != nil {
			panic(err)
		}

	}
	/*
		w.Header().Set("Content-type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		studentId := School.GetStudents()

		if err := json.NewEncoder(w).Encode(studentId); err != nil {
			panic(err)
		}
	*/
}

func TeacherIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	teachersId := School.GetTeachers()

	if err := json.NewEncoder(w).Encode(teachersId); err != nil {
		panic(err)
	}
}

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
