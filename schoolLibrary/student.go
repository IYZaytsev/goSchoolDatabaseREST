package schoollib

import (
	"fmt"
	"strconv"
)

//Student struct holds information about a student
type Student struct {
	ID        int
	ClassList Classes
	Name      string
}

//Students slice holds all the students in the school
type Students []Student

//GetStudents returns a slice of all students
func (s School) GetStudents() Students {

	allStudents := make(Students, len(s.students))

	for i := range s.students {
		allStudents[i] = s.students[i]
	}

	return allStudents

}

//SearchStudents  takes in some search parameters and returns a slice of students who match
func (s School) SearchStudents(studentName string, className string, studentID string) Students {

	//used to hold search results
	//map set is to prevent duplicate searchResults
	searchResults := Students{}
	set := make(map[int]Student)
	//converts string to int
	ID, _ := strconv.Atoi(studentID)

	//iterates over all students to find ones with IDs or Names that Match
	for i := range s.students {
		if s.students[i].ID == ID || s.students[i].Name == studentName {
			searchResults = append(searchResults, s.students[i])
			set[s.students[i].ID] = s.students[i]
		}
		//iterates over every students class list to see if the name matches
		//will not add student to slice if already in slice
		for z := range s.students[i].ClassList {
			if s.students[i].ClassList[z].Name == className {
				if _, ok := set[s.students[i].ID]; ok {
					continue
				}
				searchResults = append(searchResults, s.students[i])
			}
		}
	}

	return searchResults

}

//UpdateStudent assigns a student to a class and updates name
func (s School) UpdateStudent(studentID int, studentName string, className string) {
	//iterates over all students to find ones with IDs or Names that Match
	for i := range s.students {

		if s.students[i].ID == studentID {
			s.students[i].Name = studentName

			if className != "none" {
				for z := range s.classes {
					if s.classes[z].Name == className {
						s.students[i].ClassList = append(s.students[i].ClassList, s.classes[z])
						s.classes[z].StudentList = append(s.classes[z].StudentList, s.students[i])
					}

				}
			}
			for z := range s.classes {
				for y := range s.classes[z].StudentList {
					if s.classes[z].StudentList[y].ID == studentID {
						s.classes[z].StudentList[y].Name = studentName
					}
				}
			}

		}
	}
}

//DeleteStudent deletes student from all classes and main student list
func (s School) DeleteStudent(studentID int) {

	studentSliceALL := make(Students, 0)
	studentSliceClass := make(Students, 0)
	//iterates over all students to delete the student
	for i := range s.students {

		if s.students[i].ID == studentID {
			continue

		}
		studentSliceALL = append(studentSliceALL, s.students[i])
	}

	s.students = studentSliceALL
	fmt.Println(s.students)
	for i := range s.classes {
		for z := range s.classes[i].StudentList {
			if s.classes[i].StudentList[z].ID == studentID {
				continue
			}
			studentSliceClass = append(studentSliceClass, s.classes[i].StudentList[z])
		}
		s.classes[i].StudentList = studentSliceClass
		fmt.Println(s.classes[i].StudentList)
	}
}
