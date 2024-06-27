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

//Program's Entry Point

func main() {

}
