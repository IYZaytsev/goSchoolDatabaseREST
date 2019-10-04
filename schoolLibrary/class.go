package schoolLib

type Class struct {
    id          int
    students    Students
    teacher     Teacher
    name        string
}
type Classes []Class

func (s School) AddClass(c Class) error {
    s.classes = append(s.classes, c)
}
func(s School) findClass(id int) Class {
    for _, class in range s.classes {
        if class.id == id {
            return Class
        }
    }

    return nil
}
//func(c Class) updateNam
