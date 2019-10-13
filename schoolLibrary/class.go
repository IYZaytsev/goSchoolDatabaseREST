package schoollib

//Class struct holds information about a class
type Class struct {
	ID              int
	StudentList     Students
	AssignedTeacher Teacher
	Name            string
}

//Classes slice holds all classes in the school
type Classes []Class

func (s School) GetClasses() []int {

	classId := make([]int, len(s.classes))

	for i := range s.classes {
		classId[i] = s.classes[i].ID
	}

	return classId
}

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
