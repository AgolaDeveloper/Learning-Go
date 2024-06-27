//This mini project grades the students of a class then goes ahead to rank their performance
//... one with the highest points/marks tops the class

package main

//WE first create a composite/complex data type that'll store every detail of a student

type Students struct {
	firstName string
	lastName  string
	age       int
	subjects  []map[string]int
}

//SETTER methods that help with setting items to every student of instance STUDENTS STRUCT data type

func (s *Students) setFirstName(fname string) {
	s.firstName = fname
}

func (s *Students) setLastName(lname string) {
	s.lastName = lname
}

func (s *Students) setAge(ag int) {
	s.age = ag
}

func (s *Students) setSubjects(subjects []map[string]int) {
	s.subjects = subjects
}

//GETTER methods that help with getting/accessing items to every student of instance STUDENTS STRUCT data type

func (s Students) getFirstName() string {
	return s.firstName
}

func (s Students) getLastName() string {
	return s.lastName
}

func (s Students) getAge() int {
	return s.age
}

func (s Students) getSubjects() []map[string]int {
	return s.subjects
}

//creating a second struct CLASS that stores the students (type struct data type) ...
//... as its data elements...
//the students are stored as elements of

//Class as a storage unit of our data [students] must have a className and item/compartment that stores students
type Class struct {
	className string
	students  []Students
}

//SETTER METHODS for Class data structure

func (c *Class) setClassName(classNam string) {
	c.className = classNam
}

func (c *Class) setStudents(students []Students) {
	c.students = students
}

//GETTER METHODS for ACCESSING various data elements in our 'Class data structure'

func (c Class) getClassName() string {
	return c.className
}

func (c *Class) getStudents() []Students {
	return c.students
}

//Program's Entry Point

func main() {

}
