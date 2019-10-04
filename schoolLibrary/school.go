package schoolLib

import (
    "math/rand"
    "time"
)

type school struct {
    classes         Classes
    students        Students
    teachers        Teachers

    nextClassId     int
    nextTeacherId   int
    nextStudentId   int

    name            string
}

func NewSchool(name string) *school {
    newSchool := *new(school)

    newSchool.name = name
    newSchool.nextClassId = 10000
    newSchool.nextTeacherId = 10000
    newSchool.nextStudentId = 10000

    newSchool.autoSeed()

    return &newSchool
}

func (s school) GetNextClassId() int {
    nextClassId += 1
    return nextClassId
}
func (s school) GetNextTeacherId() int {
    nextTeacherId += 1
    return nextTeacherId
}
func (s school) GetNextStudentId() int {
    nextStudentId += 1
    return nextStudentId
}
 
func (s school) UpdateName(newName string) {
    s.name = newName
}

//hacked together auto-generate for testing purposes
func (s school) autoSeed() {
    for ; s.nextStudentId < 10120; s.nextStudentId++ {
        s.students = append(s.students, Student{id: s.nextStudentId, name: "student " + string(s.nextStudentId)})
    }
    for ; s.nextClassId < 10015; s.nextClassId++ {
        s.classes = append(s.classes, Class{id: s.nextClassId, name: "class " + string(s.nextClassId)})
    }
    for ; s.nextTeacherId < 10006; s.nextTeacherId++ {
        s.teachers = append(s.teachers, Teacher{id: s.nextStudentId, name: "teacher " + string(s.nextTeacherId)})
    }
    
    r := rand.New(rand.NewSource(time.Now().UnixNano()))

    for _, class := range s.classes {
        for _, i := range r.Perm(120)[:30] {
            class.students = append(class.students, s.students[i])
            s.students[i].classes = append(s.students[i].classes, class)
        }
        i := r.Intn(6)
        class.teacher = s.teachers[i]
        s.teachers[i].classes = append(s.teachers[i].classes, class)
    }
} 
