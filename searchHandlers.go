package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

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
	fmt.Println(parameter)
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
}

//TeacherSearch sends a slice of all teacher who match a search criteria
//if search criteria is empty returns all teachers
func TeacherSearch(w http.ResponseWriter, r *http.Request) {
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

	//if all search criteria are empty then it sends a slice of all teachers
	if len(name) == 0 && len(class) == 0 && len(id) == 0 {

		teacherSlice := School.GetTeachers()
		if err := json.NewEncoder(w).Encode(teacherSlice); err != nil {
			panic(err)
		}

	} else {
		//else it sends search results
		teacherSlice := School.SearchTeachers(name, class, id)
		if err := json.NewEncoder(w).Encode(teacherSlice); err != nil {
			panic(err)
		}

	}
}

//ClassSearch sends a slice of all classes who match a search criteria
//if search criteria is empty returns all classes
func ClassSearch(w http.ResponseWriter, r *http.Request) {
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

	//if all search criteria are empty then it sends a slice of all classes
	if len(name) == 0 && len(class) == 0 && len(id) == 0 {

		classSlice := School.GetClasses()
		if err := json.NewEncoder(w).Encode(classSlice); err != nil {
			panic(err)
		}

	} else {
		//else it sends search results
		classSlice := School.SearchClasses(name, class, id)
		if err := json.NewEncoder(w).Encode(classSlice); err != nil {
			panic(err)
		}

	}
}
