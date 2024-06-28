//practising concepts of Array DS in Go

package main

import "fmt"

//func that helps with inserting an element in an array

func arrayInsert(arr [6]int) {
	originalArray := arr
	//we Insert an element at index n
	//creating an empty array that'll contain the new elements after n is inserted
	//this is where we'll copy the merged slice upon insertion of n
	origArray := [len(arr) + 1]int{}

	var n int
	fmt.Printf("\n What Element Do you wanna Insert at index 3?\n>>: ")
	fmt.Scan(&n)

	//making original array out of the original array
	originalSlice := arr[:]

	//slicing the originalSlice, further into 2 slices;slice1 and slice2
	slice1 := originalSlice[:3]
	slice2 := originalSlice[3:]

	//appending the element to be inserted,n, on slice1 (as the last elementa t index 3)
	slice1 = append(slice1, n)

	//now merge the two slices back by appending slice2 to slice1
	slice1 = append(slice1, slice2...)

	//originalSlice = slice1
	//copy slice1 elements to an array
	origArray = [7]int(slice1)

	fmt.Println("********************************")

	fmt.Printf("Array BEFORE Insertion:\n %v\n", originalArray)
	fmt.Printf("Array AFTER Insertion:\n %v\n", origArray)

	fmt.Println("********************************")
}
func main() {
	originalArray := [6]int{20, 1200, 4, 56, 9, 100}

	//Inserting an element in ana Array
	arrayInsert(originalArray)

}
