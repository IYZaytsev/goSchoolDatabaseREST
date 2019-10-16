package schoollib

import (
	"strconv"
)

//Teacher struct holds information about a student
type Teacher struct {
	ID        int
	ClassList Classes
	Name      string
}

//Teachers slice holds all the teachers in the school
type Teachers []Teacher

//GetTeachers returns a slice of all teacher
func (s School) GetTeachers() Teachers {
	allTeacher := make(Teachers, len(s.teachers))

	for i := range s.teachers {
		allTeacher[i] = s.teachers[i]
	}

	return allTeacher
}

//SearchTeachers  takes in some search parameters and returns a slice of teachers who match
func (s School) SearchTeachers(teacherName string, className string, teacherID string) Teachers {

	//used to hold search results
	//map set is to prevent duplicate searchResults
	searchResults := Teachers{}
	set := make(map[int]Teacher)
	//converts string to int
	ID, _ := strconv.Atoi(teacherID)

	//iterates over all teachers to find ones with IDs or Names that Match
	for i := range s.teachers {
		if s.teachers[i].ID == ID || s.teachers[i].Name == teacherName {
			searchResults = append(searchResults, s.teachers[i])
			set[s.teachers[i].ID] = s.teachers[i]
		}
		//iterates over every teachers class list to see if the name matches
		//will not add teachers to slice if already in slice
		for z := range s.teachers[i].ClassList {
			if s.teachers[i].ClassList[z].Name == className {
				if _, ok := set[s.teachers[i].ID]; ok {
					continue
				}
				searchResults = append(searchResults, s.teachers[i])
			}
		}
	}

	return searchResults

}

//UpdateTeacher assigns a teacher to a class and updates name
func (s School) UpdateTeacher(teacherID int, teacherName string, className string) {
	//iterates over all teacher to find ones with IDs or Names that Match
	for i := range s.teachers {

		if s.teachers[i].ID == teacherID {
			s.teachers[i].Name = teacherName

			if className != "none" {
				for z := range s.classes {
					if s.classes[z].Name == className {
						s.teachers[i].ClassList = append(s.teachers[i].ClassList, s.classes[z])
						s.classes[z].AssignedTeacher = s.teachers[i]
					}

				}
			}
			for z := range s.classes {

				if s.classes[z].AssignedTeacher.ID == teacherID {
					s.classes[z].AssignedTeacher.Name = teacherName
				}

			}
		}
	}
}

//DeleteTeacher deletes student from all classes and main student list
func (s *School) DeleteTeacher(teacherID int) {

	teacherSliceALL := make(Teachers, 0)
	//iterates over all students to delete the student
	for i := range s.teachers {

		if s.teachers[i].ID == teacherID {
			continue

		}
		teacherSliceALL = append(teacherSliceALL, s.teachers[i])
	}

	s.teachers = teacherSliceALL

	for i := range s.classes {

		if s.classes[i].AssignedTeacher.ID == teacherID {
			s.classes[i].AssignedTeacher = Teacher{}
		}

	}

}
