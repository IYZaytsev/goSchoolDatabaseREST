package schoollib

//Student struct holds information about a student
type Student struct {
	ID        int
	ClassList Classes
	Name      string
}

//Students slice holds all the students in the school
type Students []Student

//GetStudents returns a slice of all students
func (s School) GetStudents() Student {
	/*
			studentIds := make([]int, len(s.students))

			for i := range s.students {
				studentIds[i] = s.students[i].id
		    }

		    return studentIds

	*/
	return s.students[0]
}
