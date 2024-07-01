//This mini project grades the students of a class then goes ahead to rank their performance.
//... one with the highest points/marks tops the class.

package main

import "fmt"

//WE first create a composite/complex data type that'll store every detail of a student

type Students struct {
	firstName       string
	lastName        string
	admissionNumber string
	subjects        map[string]int
}

//SETTER methods that help with setting items to every student of instance STUDENTS STRUCT data type

func (s *Students) setFirstName(fname string) {
	s.firstName = fname
}

func (s *Students) setLastName(lname string) {
	s.lastName = lname
}

func (s *Students) setAdmission(admissionNumber string) {
	s.admissionNumber = admissionNumber
}

func (s *Students) setSubjects(subjects map[string]int) {
	s.subjects = subjects
}

//GETTER methods that help with getting/accessing items to every student of instance STUDENTS STRUCT data type

func (s Students) getFirstName() string {
	return s.firstName
}

func (s Students) getLastName() string {
	return s.lastName
}

func (s Students) getAdmission() string {
	return s.admissionNumber
}

func (s Students) getSubjects() map[string]int {
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

func (c Class) getStudents() []Students {
	return c.students
}

//func for displaying all the students in the class

func (c Class) displayStudents() {
	fmt.Printf("\n%v STUDENTS : \n", c.getClassName())

	fmt.Println("####_________#_________#__________#_________#________#______________####")

	for _, value := range c.students {
		fmt.Println(value.admissionNumber)
		fmt.Printf("\n %v %v\n", value.lastName, value.firstName)

		//then range every student's map of subjects while displaying them
		//loop through the subjects-map of every value/student...
		fmt.Println("**___________________________________________________________**")
		for key, value := range value.subjects {
			//...and print every subject with its respective score

			fmt.Printf("%v: %v\n", key, value)
		}

	}

	fmt.Println("####_________#_________#__________#_________#________#______________####")

}

//Program's Entry Point

func main() {

	//U/I FEATURES FOR OUR SKOOLIE APP
	//a welcoming page for users
	welkam()

	//DEFINING LOCAL VARIABLES FOR OUR composite Data Structure

	//beginning with local variables for Students STRUCT

	//...then we move to defining local variables for Class Struct
	var className string
	students := make([]Students, 0)

	//creating instances of all our defined Data structs
	//begin with instance of Students

	//then instance of Class
	class := Class{}

	//we'll now begin setting items to our Class object
	//it should have its name
	//...and a slice of Student objects, as its data elements

	fmt.Println("Welcome to our Skoolie App")
	//user has some options that help in navigating this application's features
	//__ 1. Record Students Details/Grades
	//__ 2.

	fmt.Printf("Record Students' Details/Grades\n STUDENTS FROM YOUR CLASS ONLY\n")
	fmt.Println("Enter Class [Name of Your Class]: ")
	fmt.Scan(&className)
	class.setClassName(className)

	//before we set/store students, we gotta set values/items of Student, for every student we haave in this class

	var numOfStudents int
	fmt.Println("How many students do you wanna record their details?")
	fmt.Scan(&numOfStudents)

	//now we're going to set details for the n-number of students fed in by the teachers/user above

	for i := 0; i < numOfStudents; i++ {
		student := Students{}

		var firstName string
		var lastName string
		var admissionNumber string
		subjects := make(map[string]int)
		//	fmt.Printf("You Ready to Record details for STUDENT %v? \n CONTINUE>>\n", (1 + i))

		fmt.Println("Enter Student's First Name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter Student's Last Name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter Student's Admission Number: ")
		fmt.Scan(&admissionNumber)

		//for every student, you'll have to record subjects and marks scored
		//teacher/user must map every subject to a respective mark scored by the student

		//but how many subjects Did the students tackle?

		var numOfSubjects int

		fmt.Printf("\nHow Many Subjects for %v's Records/Ranking? \n >>: \n", student.getLastName())
		fmt.Scan(&numOfSubjects)

		var subject string
		var marks int

		for i := 0; i < numOfSubjects; i++ {
			fmt.Printf("Enter %v's Subject-%v: ", student.lastName, i+1)
			fmt.Scan(&subject)

			fmt.Printf("Enter %v's Mark-Score: ", subject)
			fmt.Scan(&marks)

			//every iteration...recording of a subject and respective mark-Score must be mapped
			subjects[subject] = marks

			//once it's mapped we need to set IT TO THE struct Student's subject item
			//student.setSubjects(subjects)
		}
		//then set marks the student has scored to our data struct of Students

		student.setLastName(lastName)
		student.setAdmission(admissionNumber)
		student.setFirstName(firstName)

		student.setSubjects(subjects)

		//then we store every instance of Student [as object element]in the Class-slice of object Students of the class

		//thus we'll append OBJECT STUDENT to the Class's local slice of Students in every iteration
		//...[every iteration sets an instance/object of Students completely]

		students = append(students, student)
		class.setStudents(students)

		//we, just appended [1] object of type Student to Struct Class's slice of struct-objects
		//... to store it to the struct we need to set it

	}

	////////////////////////////////////////////////////////////////////////////////

	//this section helps to access data elements inn both the 2 Structs
	//STARTING WITH DATA ELEMENETS STORED IN struct Class

	//__1. let's access name of our class
	classNam := class.getClassName()

	fmt.Printf("Class-Name: %v \n", classNam)

	//__2. let's access students in this class
	//classStudents := class.getStudents()
	class.displayStudents()

	//fmt.Println("Welcome to our SKoolie App")

}

//defining our welcoming function that welcomes displays a welcoming message upon openning the Application
func welkam() {
	fmt.Println("WELCOME TO SKOOLIE APP")
	fmt.Println("++++++++++++++++++++++++++++++++")

}
