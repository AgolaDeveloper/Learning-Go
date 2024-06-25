package main

import "fmt"

func advancedArray() {
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

	/*if count == 0 {
		fmt.Printf("Element %v NOT available in Our Array \n \n", element)

	}*/
}

func main() {
	LearningGo := "My First Commit:"

	fmt.Printf("I'm passionate about Go. Just getting started \n Here's %v\n", LearningGo)

	fmt.Println()

	//Go-Slices
	//MySlice()
	advancedArray()

}
