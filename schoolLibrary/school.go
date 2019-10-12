package schoollib

import (
	"fmt"
)

//School struct contains information about classes, students, and teachers.
type School struct {
	classes  Classes
	students Students
	teachers Teachers

	nextClassID   int
	nextTeacherID int
	nextStudentID int

	name string
}

//NewSchool initializes a school
func NewSchool(name string) *School {

	newSchool := *new(School)

	newSchool.name = name
	newSchool.nextClassID = 10000
	newSchool.nextTeacherID = 10000
	newSchool.nextStudentID = 10000

	//newSchool.autoSeed()

	return &newSchool
}

//AddStudent adds a new student to the school
func (s *School) AddStudent(studentName string) Student {
	newStudent := Student{ID: s.nextStudentID, Name: studentName}
	s.students = append(s.students, Student{ID: s.nextStudentID, Name: studentName})
	s.nextStudentID++
	fmt.Println(s.nextStudentID)
	return newStudent

}

//hacked together auto-generate for testing purposes
/*
func (s *School) autoSeed() {
	for ; s.nextStudentId < 10120; s.nextStudentId++ {
		s.students = append(s.students, Student{ID: s.nextStudentId, Name: fmt.Sprintf("student %d", s.nextStudentId)})
	}
	for ; s.nextClassId < 10015; s.nextClassId++ {
		s.classes = append(s.classes, Class{id: s.nextClassId, name: fmt.Sprintf("class %d", s.nextClassId)})
	}
	for ; s.nextTeacherId < 10006; s.nextTeacherId++ {
		s.teachers = append(s.teachers, Teacher{id: s.nextTeacherId, Name: fmt.Sprintf("teacher %d", s.nextTeacherId)})
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for z := range s.classes {
		for i := range r.Perm(120)[:30] {
			s.classes[z].students = append(s.classes[z].students, s.students[i])
			s.students[i].ClassList = append(s.students[i].ClassList, s.classes[z])
		}
		i := r.Intn(6)
		s.classes[z].teacher = s.teachers[i]
		s.teachers[i].classes = append(s.teachers[i].classes, s.classes[z])
	}

}
*/
