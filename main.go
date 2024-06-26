package main

import "fmt"

func searchArray() {
	//trying to do a SEARCH OPERATION on our array
	fmt.Printf("***************************************************\n**************************************************\n\n")
	fmt.Println("Beginning of our ADVANCED ARRAY SECTION")

	ourArray := [10]int{}

	//and every element in our array'd be double its previous contagious element
	//...where the first element is 4
	ourArray[0] = 2
	for i := 1; i < len(ourArray); i++ {

		ourArray[i] = ourArray[i-1] * 2
	}

	fmt.Printf("OUR ARRAY: \n      %v", ourArray)

	//NOW we gotta SEARCH for an element [128 in this case]

	element := 1
	count := 0
	for i := 0; i < len(ourArray); i++ {

		if ourArray[i] == element {
			fmt.Printf("Element %v Available in Our Array \n It's stored at index %v \n", element, i)
			count++
			//elemPresent = true
		} else if i == len(ourArray)-1 && count == 0 {
			fmt.Printf("\nElement %v NOT available in Our Array \n \n", element)

		}
	}

	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
}

//Practising COMPARING ELEMENTS in an Array

func compareElements() {
	fmt.Println("//Practising COMPARING ELEMENTS in an Array//")
	array := [10]int{30, 20, 100, 10, 300, 1000, 25, 6, 1, 35}

	fmt.Printf("\nOur Array before sorting:\n %v\n", array)

	//now we COMPARE the subsequent elements

	for i := 0; i < len(array)-1; i++ {

		if array[i] < array[i+1] {
			fmt.Printf("%v is greater than %v\n", array[i+1], array[i])

		} else {
			fmt.Printf("%v is lesser than %v\n", array[i+1], array[i])

		}
	}

	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
}

//SWAPPING THE FIRST 2 ELEMENTS OF OUR ARRAY

func swapElements() {
	array := [10]int{30, 20, 100, 10, 300, 1000, 25, 6, 1, 35}

	fmt.Printf("\nOUR ARRAY BEFORE SWAPPING FIRST 2 ELEMENTS:\n %v\n", array)
	fmt.Printf("Original Positions of the Elements:\n Index 0: %v \n Index 1: %v \n", array[0], array[1])

	//we'll have to create a temporary variable, TEMP
	//...TEMP shall help us with swapping 2 elements
	var temp int

	temp = array[0]
	array[0] = array[1]
	array[1] = temp

	fmt.Printf("\nOUR ARRAY AFTER SWAPPING FIRST 2 ELEMENTS:\n %v\n", array)
	fmt.Printf("NEW Positions of the Elements:\n Index 0: %v \n Index 1: %v \n", array[0], array[1])

	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
}

//SORTING Array ELEMENTS in Descending Order

func sortDescend() {
	array := [10]int{30, 20, 100, 10, 300, 1000, 25, 6, 1, 35}

	fmt.Printf("Our Array before Sorting:\n %v\n", array)

	//comparing every element against all elements in the array

	for i := 0; i < len(array)-1; i++ {

		var temp int

		//we first assume current element we're doing comparison against is the largest
		largest := array[i]

		//comparison begins from the immediate subsequent element of the current [largest] element
		for j := (i + 1); j < len(array); j++ {

			if array[j] > largest {
				largest = array[j]

				//swapping new largest value [found in inner loop] with our initial-largest i [of outer loop]

				temp = array[i]
				array[i] = largest
				array[j] = temp
			}
		}
	}

	fmt.Printf("Assorted Array [Descending]:\n %v\n", array)

	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
}

//SORTING Array ELEMENTS in Descending Order
func sortAscend() {
	array := [10]int{30, 20, 100, 10, 300, 1000, 25, 6, 1, 35}

	fmt.Printf("Array before Sort:\n %v\n", array)

	//create a TEMPORARY variable that will help us with swapping two number upon comparison
	var temp int

	for i := 0; i < len(array)-1; i++ {
		//every iteration of our outter loop...
		//...we set array[i] current element as the smallest
		//...and compare it against the rest of the subsequent numbers

		smallest := array[i]

		for j := (i + 1); j < len(array); j++ {
			//comparing all the elemets in inner loop against current smallest of every iteration

			if array[j] < smallest {
				//then set array[j] as the smallest
				smallest = array[j]

				temp = array[i]
				array[i] = smallest
				array[j] = temp
			}
		}
	}

	fmt.Printf("Assorted Array [Ascending]:\n %v\n", array)
	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
}

//FINDING THE LARGEST ELEMENT IN THE ARRAY
func laGest() {
	fmt.Println("//FINDING THE LARGEST ELEMENT IN AN ARRAY//")

	//first step to the problem is sorting the array
	//either in ascending or descending order

	//In Descending Order, first element in the array is the largest...
	//while last element in the array is the largest in Ascending Order.

	//sorting in Descending Order, in this case
	array := [10]int{30, 20, 100, 10, 300, 1000, 25, 6, 1, 35}

	var temp int

	for i := 0; i < len(array)-1; i++ {
		largest := array[i]

		for j := (i + 1); j < len(array); j++ {
			if array[j] > largest {
				largest = array[j]

				temp = array[i]
				array[i] = largest
				array[j] = temp
			}
		}
	}

	//Upon Sorting in Descending Order, Largest element is stored at the first index
	fmt.Printf("\nLargest Number: \n >> %v\n", array[0])

	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
}

//FINDING THE LARGEST ELEMENT IN THE ARRAY
func smallestElement() {
	fmt.Println("//FINDING THE SMALLEST ELEMENT IN AN ARRAY//")

	//first step to the problem is sorting the array
	//either in ascending or descending order

	//In Descending Order, last element in the array is the smallest...
	//while first element in the array is the smallest in Ascending Order.

	//sorting in Descending Order, in this case
	array := [10]int{30, 20, 100, 10, 300, 1000, 25, 6, 1, 35}

	var temp int

	for i := 0; i < len(array)-1; i++ {
		largest := array[i]

		for j := (i + 1); j < len(array); j++ {
			if array[j] > largest {
				largest = array[j]

				temp = array[i]
				array[i] = largest
				array[j] = temp
			}
		}
	}

	//Upon Sorting in Descending Order, Largest element is stored at the first index
	fmt.Printf("\nSmallest Number: \n >> %v\n", array[len(array)-1])

	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
}

//The Program's ENTRY POINT

func main() {
	LearningGo := "My First Commit:"

	fmt.Printf("I'm passionate about Go. Just getting started \n Here's %v\n", LearningGo)

	fmt.Println()

	//Go-Slices
	//MySlice()
	searchArray()
	compareElements()
	swapElements()
	sortDescend()
	sortAscend()
	laGest()
	smallestElement()
}
