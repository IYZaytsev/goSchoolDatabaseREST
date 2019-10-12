package schoollib

//Class struct holds information about a class
type Class struct {
	id       int
	students Students
	teacher  Teacher
	name     string
}

//Classes slice holds all classes in the school
type Classes []Class

func (s School) GetClasses() []int {

	classId := make([]int, len(s.classes))

	for i := range s.classes {
		classId[i] = s.classes[i].id
	}

	return classId
}

func (s School) GetClasseTeacher(classId int) int {

	for i := range s.classes {
		if s.classes[i].id == classId {
			return s.classes[i].teacher.id
		}
	}

	return -1
}

func (s School) GetClasseStudents(classId int) []int {

	for i := range s.classes {
		if s.classes[i].id == classId {
			studentsID := make([]int, len(s.classes[i].students))
			for z := range s.classes[i].students {
				studentsID[z] = s.classes[i].students[z].ID
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
