package schoollib

import "strconv"

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
