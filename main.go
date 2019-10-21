package main

import (
	"fmt"
	"log"
	"net/http"

	schoollib "./schoolLibrary"
)

//School is captial here so it can be accesed within the handlers.go file
var School *schoollib.School

func main() {

	School = schoollib.NewSchool("oak grove")
	//ADDING STUDENTS
	School.AddStudent("Colby Legge")
	School.AddStudent("Mariya Snider")
	School.AddStudent("Derek Hampton")
	School.AddStudent("Usamah Clemons")
	School.AddStudent("Elspeth Morgan")
	School.AddStudent("Pranav Macfarlane")
	School.AddStudent("Neel Mcneill")
	School.AddStudent("Dayna Barr")
	School.AddStudent("Wil Goldsmith")
	School.AddStudent("Ishmael Mansell")
	School.AddStudent("Shola Hebert")
	School.AddStudent("Sumaiyah Person")
	School.AddStudent("Timur Beasley")
	School.AddStudent("Lance Mccullough")

	//ADDING TEACHERS
	School.AddTeacher("Ritchie Clarke")
	School.AddTeacher("Gia Leech")
	School.AddTeacher("Jay-Jay Dawe")
	School.AddTeacher("Usmaan Hall")
	School.AddTeacher("Kaan Goodwin")
	School.AddTeacher("Yassin Cook")

	//ADDING CLASSES
	School.AddClass("American Literature")
	School.AddClass("British Literature")
	School.AddClass("World Literature")
	School.AddClass("Physical Education (P.E.)")
	School.AddClass("Physics")
	School.AddClass("Meteorology")
	School.AddClass("Precalculus")
	School.AddClass("Statistics")
	School.AddClass("Trigonometry")
	School.AddClass("Geometry")

	//Adding relationships
	School.UpdateStudent(10000, "Colby Legge", "Geometry")
	School.UpdateStudent(10000, "Colby Legge", "Trigonometry")
	School.UpdateStudent(10000, "Colby Legge", "Statistics")
	School.UpdateStudent(10000, "Colby Legge", "Precalculus")
	School.UpdateStudent(10001, "Mariya Snider", "Meteorology")
	School.UpdateStudent(10001, "Mariya Snider", "Geometry")
	School.UpdateStudent(10003, "Usamah Clemons", "British Literature")
	School.UpdateStudent(10003, "Usamah Clemons	", "Geometry")
	School.UpdateStudent(10006, "Neel Mcneill", "Geometry")
	School.UpdateStudent(10006, "Neel Mcneill", "World Literature")
	School.UpdateStudent(10009, "Ishmael Mansell", "World Literature")
	School.UpdateStudent(10009, "Ishmael Mansell", "Geometry")
	School.UpdateStudent(10010, "Shola Hebert", "British Literature")
	School.UpdateStudent(10010, "Shola Hebert", "Physical Education (P.E.)")

	School.UpdateTeacher(10000, "Ritchie Clarke", "British Literature")
	School.UpdateTeacher(10000, "Ritchie Clarke", "World Literature")
	School.UpdateTeacher(10001, "Gia Leech", "Physical Education (P.E.)")
	School.UpdateTeacher(10002, "Jay-Jay Dawe", "Geometry")
	School.UpdateTeacher(10003, "Usmaan Hall", "Statistics")
	School.UpdateTeacher(10004, "Kaan Goodwin", "Precalculus")
	School.UpdateTeacher(10005, "Yassin Cook", "Trigonometry")

	fmt.Println("starting router")

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
