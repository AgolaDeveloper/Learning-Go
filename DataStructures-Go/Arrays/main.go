//practising concepts of Array DS in Go

package main

import "fmt"

//func that helps with inserting an element in an array

/*func arrayInsert() {
	originalArray := [6]int{20, 1200, 4, 56, 9, 100}
	copyArray := [6]int{}
	copyArray = originalArray

	//	originalArray := arr
	//we Insert an element at index n
	//creating an empty array that'll contain the new elements after n is inserted
	//this is where we'll copy the merged slice upon insertion of n
	//origArray := [len(arr) + 1]int{}

	var n int
	fmt.Printf("\n What Element Do you wanna Insert at index 3?\n>>: ")
	fmt.Scan(&n)

	//making original array out of the original array
	originalSlice := copyArray[:]

	//slicing the originalSlice, further into 2 slices;slice1 and slice2
	slice1 := originalSlice[:3]
	slice2 := originalSlice[3:]
	copySlice2 := slice2
	//appending the element to be inserted,n, on slice1 (as the last elementa t index 3)

	copySlice1 := slice1
	copySlice1 = append(copySlice1, n)
	//slice1 = append(slice1, n)
	//copySlice1 := make([]int, 0)

	//now merge the two slices back by appending slice2 to slice1
	//slice1 = append(slice1, slice2...)
	copySlice1 = append(copySlice1, copySlice2...)
	//originalSlice = slice1
	//copy slice1 elements to an array
	//origArray = [7]int(slice1)

	fmt.Println("********************************")

	fmt.Printf("Array BEFORE Insertion:\n %v\n", originalArray)
	fmt.Printf("Array AFTER Insertion:\n %v\n", copySlice1)

	fmt.Println("********************************")
}*/

func insert() {
	array := [5]int{3, 10, 6, 11, 50}
	slice := make([]int, len(array)+1)

	//fmt.Println("SLICE: ", slice)

	slice = array[:]

	//fmt.Println("SLICE: ", slice)

	//slice := array[:]
	slice = append(slice, 0)
	//fmt.Println("SLICE: ", slice)

	var element int
	fmt.Println("What do you wanna Insert at index 2?: ")
	fmt.Scan(&element)
	//inserting at index 2
	for i := 4; i > 1; i-- {
		slice[i+1] = slice[i]

	}
	slice[2] = element

	fmt.Println("Array BEFORE Insertion: ", array)
	fmt.Println("Array AFTER Insertion: ", slice)

}
func main() {
	//	originalArray := [6]int{20, 1200, 4, 56, 9, 100}

	//Inserting an element in ana Array
	//arrayInsert()
	insert()
}
