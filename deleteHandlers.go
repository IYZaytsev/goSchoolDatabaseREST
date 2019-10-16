package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

//StudentDelete Deletes Student
func StudentDelete(w http.ResponseWriter, r *http.Request) {

	//gets passed ID for deletion
	body, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	parameter := string(body)

	studentID, _ := strconv.Atoi(parameter)
	School.DeleteStudent(studentID)
}

//TeacherDelete Deletes teacher
func TeacherDelete(w http.ResponseWriter, r *http.Request) {

	//gets passed ID for deletion
	body, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	parameter := string(body)

	teacherID, _ := strconv.Atoi(parameter)
	School.DeleteTeacher(teacherID)
}

//TeacherDelete Deletes teacher
func ClassDelete(w http.ResponseWriter, r *http.Request) {

	//gets passed ID for deletion
	body, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	parameter := string(body)

	classID, _ := strconv.Atoi(parameter)
	School.DeleteClass(classID)
}
