package schoolLib

import (
    "math/rand"
    "time"
    "encoding/json"

    "fmt"
)

type School struct {
    classes         Classes
    students        Students
    teachers        Teachers

    nextClassId     int
    nextTeacherId   int
    nextStudentId   int

    name            string
}

func NewSchool(name string) School {
    newSchool := *new(School)

    newSchool.name = name
    newSchool.nextClassId = 10000
    newSchool.nextTeacherId = 10000
    newSchool.nextStudentId = 10000

    newSchool.autoSeed()

    return newSchool
}

func (s School) GetNextClassId() int {
    s.nextClassId += 1
    return s.nextClassId
}
func (s School) GetNextTeacherId() int {
    s.nextTeacherId += 1
    return s.nextTeacherId
}
func (s School) GetNextStudentId() int {
    s.nextStudentId += 1
    return s.nextStudentId
}
 
func (s School) UpdateName(newName string) {
    s.name = newName
}

//hacked together auto-generate for testing purposes
func (s School) autoSeed() {
    for ; s.nextStudentId < 10120; s.nextStudentId++ {
        s.students = append(s.students, Student{id: s.nextStudentId, name: fmt.Sprintf("student %d", s.nextStudentId)})
    }
    for ; s.nextClassId < 10015; s.nextClassId++ {
        s.classes = append(s.classes, Class{id: s.nextClassId, name: fmt.Sprintf("class %d", s.nextClassId)})
    }
    for ; s.nextTeacherId < 10006; s.nextTeacherId += 1 {
        s.teachers = append(s.teachers, Teacher{id: s.nextTeacherId, name: fmt.Sprintf("teacher %d", s.nextTeacherId)})
    }
    
    r := rand.New(rand.NewSource(time.Now().UnixNano()))

    for _, class := range s.classes {
        for i, _ := range r.Perm(120)[:30] {
            class.students = append(class.students, s.students[i])
            s.students[i].classes = append(s.students[i].classes, class)
        }
        i := r.Intn(6)
        class.teacher = s.teachers[i]
        s.teachers[i].classes = append(s.teachers[i].classes, class)
    }
} 
