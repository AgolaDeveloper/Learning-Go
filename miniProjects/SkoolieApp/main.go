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

/*func (s Students) getFirstName() string {
	return s.firstName
}*/

func (s Students) getLastName() string {
	return s.lastName
}

/*func (s Students) getAdmission() string {
	return s.admissionNumber
}

func (s Students) getSubjects() map[string]int {
	return s.subjects
}*/

//creating a second struct CLASS that stores the students (type struct data type) ...
//... as its data elements...
//the students are stored as elements of

//Class as a storage unit of our data [students] must have a className and item/compartment that stores students
type Class struct {
	className string
	students  []Students
}

//SETTER METHODS for Class data structure

/*func (c *Class) setClassName(classNam string) {
	c.className = classNam
}*/

func (c *Class) setStudents(students []Students) {
	c.students = students
}

//GETTER METHODS for ACCESSING various data elements in our 'Class data structure'

func (c Class) getClassName() string {
	return c.className
}

/*func (c Class) getStudents() []Students {
	return c.students
}*/

////////////////////////////////////////////////////////////////////////////////////////////////////////
//RETRIEVAL METHODS: helping with retrieving INFO about the STUDENTS

//Every RETRIEVAL METHOD MUST FIRST CHECK WHETHER OUR DATA STRUCTURE of STUDENT OBJECTS is Populated
//...else it returns "NO RECORDS FOR THE CLASS"

//func to help with checking if the Class DATA STRUCT is populated
func (c Class) classEmpty() bool {
	//check if slice of object-Students in the Main Data Struct is empty

	if c.students == nil {
		return true
	} else {
		return false
	}
}

//func for displaying all the students in the class

func (c Class) displayStudents() {

	//we only disply STUDENT(S) Details if the item (slice of objects) storing students'details...
	//...in Class data struct isn't empty

	if c.classEmpty() {
		fmt.Println("Ooops...! No Student Found")
		fmt.Println("Class Teacher hasn't Recorded any Student Details for the class")

	} else {
		//else, if the class is not empty
		//display the Student's available

		fmt.Printf("\n%v STUDENTS : \n", c.getClassName())

		fmt.Println("####_________#_________#__________#_________#________#______________####")

		for _, value := range c.students {
			fmt.Printf("\n\nAdmission Number: %v", value.admissionNumber)
			fmt.Printf("\n %v %v\n", value.lastName, value.firstName)

			//then range every student's map of subjects while displaying them
			//loop through the subjects-map of every value/student...
			fmt.Println("*___________________________________________________________*")

			for key, value := range value.subjects {
				//...and print every subject with its respective score

				fmt.Printf("%v: %v\n", key, value)

			}

		}

		fmt.Println("####_________#____________________________________#_________####")

	}
}

//Adding method that RETURNS a map of total marks scored by every student
//and then store it in a new map

func (c Class) totalMarks() map[string]int {

	totalMrks := make(map[string]int)

	if c.classEmpty() {
		fmt.Println("Ooops... No Students!")
	} else {
		//loop through slice containing Student objects, in the class struct

		for _, value := range c.students {
			//...then loop through every students map of subjects
			//... while summing the total marks for every student

			total := 0
			//count := 0
			for _, v := range value.subjects {
				total += v

			}
			//for every student's total; map it to the student's name

			//concatenating student's last and first names
			name := value.lastName + " " + value.firstName

			totalMrks[name] = total
		}
	}

	return totalMrks
}

//this methods prints/displays total marks scored for every student in the class

func (c Class) printTotalMarks() {
	//get value [a map] returned by the Struct's totalMarks() method
	//... then loop while printing every student's total marks

	//... fmt.Printf("\nSTUDENT NAME %-20v TOTAL MARKS\n")
	for key, value := range c.totalMarks() {
		fmt.Printf(" %-20v : %v\n", key, value)

	}
}

//func that returns details of a specific student
//...given an admission number

func (c Class) studentMarks(admissionNumber string) {
	//first check if the admission number entered by user exist

	for _, value := range c.students {
		if value.admissionNumber == admissionNumber {
			//then print the respective student's details
			name := value.lastName + " " + value.firstName

			fmt.Printf("\nAdmission Number: %v\n", value.admissionNumber)
			fmt.Printf("Name: %v\n", name)

			//then range through the subjetc's map and find totalmarks
			total := 0
			for key, value := range value.subjects {
				total += value

				fmt.Printf("%v: %v\n", key, value)
			}

			fmt.Printf("TOTAL MARKS: %v", total)

		} else {
			fmt.Printf("\nOoops...! Student with admission %v doesn't Exist\n", admissionNumber)
		}
	}
}

//Program's Entry Point

func main() {

	//U/I FEATURES FOR OUR SKOOLIE APP
	//a welcoming page for users
	welkam()

	//...then we move to defining local variables for Class Struct
	//var className string
	students := make([]Students, 0)

	//creating instances of all our defined Data structs
	//begin with instance of Students

	//then instance of Class
	class := Class{className: "Form4_East"}

	for {
		//IT'S THE USER'S TO CHOOSE WHAT THEY WANT

	choose:
		var userChoice int
		fmt.Println("CHOOSE: >>")
		fmt.Println("1. RECORD STUDENTS [Details]")
		fmt.Println("2. RETRIEVE STUDENTS [Details]")

		fmt.Scan(&userChoice)

		switch userChoice {
		case 1:
			//DEFINING LOCAL VARIABLES FOR OUR composite Data Structure

			//beginning with local variables for Students STRUCT

			//we'll now begin setting items to our Class object
			//it should have its name
			//...and a slice of Student objects, as its data elements

			fmt.Println("Welcome to our Skoolie App")
			//user has some options that help in navigating this application's features
			//__ 1. Record Students Details/Grades
			//__ 2.

			fmt.Printf("Record Students' Details/Grades\n STUDENTS FROM %v ONLY\n", class.className)

			/*fmt.Println("Enter Class [Name of Your Class]: ")
			fmt.Scan(&className)
			class.setClassName(className)*/

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

			//break

		case 2:

			//once you take the CHOICE to RETRIEVE THE STUDENTS' DETAIL
			//... YOU AGAIN GOTTA CHOOSE[BE SPECIFIC WITH THE EXACT OPERATION]
			var chus int
			fmt.Println("CHOOSE: ")
			fmt.Println("1. STUDENT'S DETAILS")
			fmt.Println("2. STUDENTS' TOTAL MARKS")

			fmt.Scan(&chus)
			switch chus {
			case 1:
				class.displayStudents()

			case 2:
				class.printTotalMarks()

			}

		default:
			fmt.Println("INVALID CHOICE RETRY...")
			goto choose
		}

		////////////////////////////////////////////////////////////////////////////////

	yourChoice:
		var cont string

		fmt.Println("y. Continue")
		fmt.Println("n. EXIT ")

		fmt.Scan(&cont)

		switch cont {
		case "y":
			fmt.Println("Hurray Continue Exploring SKOOLIE App")
			continue
		case "n":
			fmt.Println("Ooops...! Sad to See You Leave our Cool SKOOLIE App")
			fmt.Println("CHEERS")
			return
		default:
			goto yourChoice
		}

	}

}

//defining our welcoming function that welcomes displays a welcoming message upon openning the Application
func welkam() {
	fmt.Println("WELCOME TO SKOOLIE APP")
	fmt.Println("++++++++++++++++++++++++++++++++")

}
