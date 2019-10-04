package schoolLib

type Class struct {
    id          int
    students    Students
    teacher     Teacher
    name        string
}
type Classes []Class

func (s School) GetClasses() Classes {
    return s.classes
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
