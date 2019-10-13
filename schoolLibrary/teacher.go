package schoollib

import "strconv"

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
