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
		for j := (i + 1); j < len(array); i++ {

			if array[j] > largest {
				largest = array[j]

			}
			temp = array[i]
			array[i] = largest
			array[j] = temp

		}
		//swapping new largest with our initial-largest i

	}

	fmt.Printf("Assorted Array [Descending]:\n %v\n", array)

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

}
