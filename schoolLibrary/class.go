package schoollib

import "strconv"

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

/*
func (s School) GetClasseTeacher(classId int) int {

	for i := range s.classes {
		if s.classes[i].ID == classId {
			return s.classes[i].AssignedTeacher.ID
		}
	}

	return -1
}

func (s School) GetClasseStudents(classId int) []int {

	for i := range s.classes {
		if s.classes[i].ID == classId {
			studentsID := make([]int, len(s.classes[i].StudentList))
			for z := range s.classes[i].StudentList {
				studentsID[z] = s.classes[i].StudentList[z].ID
			}
			return studentsID
		}
	}
	return nil
}
*/
/*

func (s school) AddClass(c Class) error {
    s.classes = append(s.classes, c)
}
func(s school) findClass(id int) Class {
    for _, class := range s.classes {
        if class.id == id {
            return Class
        }
    }
}

*/
