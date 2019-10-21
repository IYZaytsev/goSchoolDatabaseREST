package schoollib

import (
	"strconv"
)

//Class struct holds information about a class
type Class struct {
	ID              int
	StudentList     Students
	AssignedTeacher Teacher
	Name            string
}

//Classes slice holds all classes in the school
type Classes []Class

//GetClasses returns a slice all classes
func (s School) GetClasses() Classes {
	allClasses := make(Classes, len(s.classes))

	for i := range s.classes {
		allClasses[i] = s.classes[i]
	}

	return allClasses
}

//SearchClasses  takes in some search parameters and returns a slice of Classes who match
func (s School) SearchClasses(className string, teacherName string, classID string) Classes {

	//used to hold search results
	searchResults := Classes{}
	//converts string to int
	ID, _ := strconv.Atoi(classID)

	//iterates over all classes to find ones with IDs or Names that Match
	for i := range s.classes {
		if s.classes[i].ID == ID || s.classes[i].Name == className || (s.classes[i].AssignedTeacher.Name == teacherName && len(teacherName) > 0) {
			searchResults = append(searchResults, s.classes[i])

		}

	}

	return searchResults

}

//UpdateClass updates a class name
func (s School) UpdateClass(classID int, className string) {
	//iterates over all class to find ones with IDs or Names that Match
	for i := range s.classes {

		if s.classes[i].ID == classID {
			s.classes[i].Name = className
			for z := range s.students {
				for y := range s.students[z].ClassList {
					if s.students[z].ClassList[y].ID == classID {
						s.students[z].ClassList[y].Name = className
					}
				}
			}
			for z := range s.teachers {
				for y := range s.teachers[z].ClassList {
					if s.teachers[z].ClassList[y].ID == classID {
						s.teachers[z].ClassList[y].Name = className
					}
				}
			}
		}
	}
}

//DeleteClass deletes class
func (s *School) DeleteClass(classID int) {

	classSliceALL := make(Classes, 0)
	classSliceStudent := make(Classes, 0)
	classSliceTeacher := make(Classes, 0)
	//iterates over all students to delete the student
	for i := range s.classes {

		if s.classes[i].ID == classID {
			continue

		}
		classSliceALL = append(classSliceALL, s.classes[i])
	}

	s.classes = classSliceALL

	for i := range s.students {

		for z := range s.students[i].ClassList {
			if s.students[i].ClassList[z].ID == classID {
				continue
			}
			classSliceStudent = append(classSliceStudent, s.students[i].ClassList[z])
		}
		s.students[i].ClassList = classSliceStudent
		classSliceStudent = nil
	}

	for i := range s.teachers {
		for z := range s.teachers[i].ClassList {
			if s.teachers[i].ClassList[z].ID == classID {
				continue
			}
			classSliceTeacher = append(classSliceStudent, s.teachers[i].ClassList[z])
		}
		s.teachers[i].ClassList = classSliceTeacher
		classSliceTeacher = nil
	}
}
