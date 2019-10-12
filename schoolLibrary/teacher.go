package schoollib

import "fmt"

//Teacher struct holds information about a student
type Teacher struct {
	id      int
	classes Classes
	Name    string
}

//Teachers slice holds all the teachers in the school
type Teachers []Teacher

//GetTeachers returns a slice of all teachers Id
func (s School) GetTeachers() []int {
	teacherIds := make([]int, len(s.teachers))

	for i := range s.teachers {
		teacherIds[i] = s.teachers[i].id
	}
	fmt.Println(teacherIds)
	return teacherIds
}
