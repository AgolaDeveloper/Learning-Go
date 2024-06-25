// ADVANCING ARRAY OPERATIONS
package main

import "fmt"

func advancingArray() {
	//trying to do a SEARCH OPERATION on our array

	ourArray := [10]int{}

	//and every element in our array'd be square of its previous contagious element
	//...where the first element is 4
	ourArray[0] = 2
	for i := 1; i < len(ourArray); i++ {
		ourArray[i] = ourArray[i-1] * ourArray[i-1]
	}

	fmt.Println("OUR ARRAY: ", ourArray)
}
