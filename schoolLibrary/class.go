package schoolLib

type Class struct {
    id          int
    students    Students
    teacher     Teacher
    name        string
}
type Classes []Class
