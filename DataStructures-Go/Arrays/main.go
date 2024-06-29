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

//1st METHOD OF INSERTION; BY SHIFTING
func insert1(array [6]int) {

	fmt.Println("INSERTION BY SHIFTING")

	//array := [6]int{20, 1200, 4, 56, 9, 100}

	//making a slice that'll copy the original array to...
	//slice := make([]int, 0)

	//coping our array into the slice
	//slice is dynamic; allows us to shift elements

	slice := array[:]
	slice = append(slice, 0)

	var element int
	fmt.Printf("What Element Do You Wanna Insert?\n >>>:  \n")
	fmt.Scan(&element)

	var index int
	fmt.Printf("\nAt What Index Do You Wanna Insert Element %v?\n >>>:  \n", element)
	fmt.Scan(&index)

	//inserting at index 2
	for i := len(array) - 1; i >= index; i-- {
		slice[i+1] = slice[i]

	}
	slice[index] = element

	fmt.Println("Array BEFORE Insertion: ", array)
	fmt.Println("Array AFTER Insertion: ", slice)

}

//2nd METHOD OF INSERTION;BY SLICING
func insert2(arr [6]int) {

	fmt.Println("INSERTION BY SLICING")
	mySlice := make([]int, 0)
	//here we just make slice outta the array and slice it further
	slice := arr[:]

	var n int
	fmt.Println("At what index do you wanna insert an element?:")
	fmt.Scan(&n)

	var e int
	fmt.Println("  element?:")
	fmt.Scan(&e)

	//slice it into two
	//first, from 0 to immediately before the index of insertion
	slice1 := slice[:n]
	//append it to mySlice
	mySlice = append(mySlice, slice1...)

	//second one, from index of insertion's element to the last element
	slice2 := slice[n:]

	//so we'll append to slice1, the element to insert and slice2 respectively

	mySlice = append(mySlice, e)
	mySlice = append(mySlice, slice2...)

	fmt.Println("Array BEFORE Insertion: ", arr)
	fmt.Println("Array AFTER Insertion: ", mySlice)
}

//DELETING AN ELEMENT
// 1. By Shifting
//adding a function that'll delete an element at a given index in our Array
func delete(array [6]int) {
	//we first create a slice of the entire array

	mySlice := array[:]
	slice := make([]int, 0)
	slice2 := make([]int, 0, 20)

	//copy mySlice to slice
	//copy(slice, mySlice)
	slice = append(slice, mySlice...)

	var index int
	fmt.Println("At what index do you wanna delete an element?: ")
	fmt.Scan(&index)

	//Now shift all the elements in our slice to the left
	//...where we begin with the immediate element to the right of indexed element; replacing it

	for i := index; i < len(slice)-1; i++ {
		slice[i] = slice[i+1]

		if slice[i] == len(slice)-2 {
			slice2 = slice[:len(slice)-1]
		}
	}

	fmt.Printf("\n Before Deletion: %v\n", array)
	fmt.Printf("\n After Deletion: %v\n", slice2)

}

//DELETION BY SLICING

func delete1(array [6]int) {
	//first make slice outta the backing array
	slice := array[:]

	//create a second slice from which we'll manipulate the elements without reflecting on the backing array
	mySlice := make([]int, 0)

	mySlice = append(mySlice, slice...)

	//now slice it into two different slices
	var index int
	fmt.Println("At what index do you wanna delete an element?: ")
	fmt.Scan(&index)

	//first slice upto b4 the index
	slice1 := mySlice[:index]

	//then append to ourSlice
	ourSlice := make([]int, 0)

	ourSlice = append(ourSlice, slice1...)
	//again slice another from index+1

	slice2 := mySlice[index+1:]

	ourSlice = append(ourSlice, slice2...)

	fmt.Printf("\n Before Deletion: %v\n", array)
	fmt.Printf("\n After Deletion: %v\n", ourSlice)

}

func main() {
	originalArray := [6]int{20, 1200, 4, 56, 9, 100}

	//Inserting an element in ana Array
	//arrayInsert()
	insert1(originalArray)
	insert2(originalArray)
	delete(originalArray)
	delete1(originalArray)

}
