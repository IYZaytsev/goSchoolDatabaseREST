package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

func StudentIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	studentId := School.GetStudents()

	if err := json.NewEncoder(w).Encode(studentId); err != nil {
		panic(err)
	}
}

func TeacherIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	teachersId := School.GetTeachers()

	if err := json.NewEncoder(w).Encode(teachersId); err != nil {
		panic(err)
	}
}

//Creates a new student and takes in
func StudentCreation(w http.ResponseWriter, r *http.Request) {

	r.FormValue("studentName")
	studentID := School.AddStudent(r.FormValue("studentName"))
	w.Header().Set("Content-type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(studentID); err != nil {
		panic(err)
	}

}

func test(w http.ResponseWriter, r *http.Request) {
	url := "http://localhost:8080/students"
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "test")
	res, getErr := myClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	student1 := schoolLib.Student{}
	jsonErr := json.Unmarshal(body, &student1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println(student1.ID)

}
