// this program file adds items to the map
package main

import "fmt"

//func that adds item to the map
//takes 2 parameters, ourStruct and isMapEmpty
func addItem(ourMap map[string]int, isMapEmpty bool) {

	var dish string
	var dishPrice int
	//add directly if map is empty
	if isMapEmpty {
		//if map is empty directly add
		//prompt user to enter key and values repectively
		fmt.Println("Enter Dish: ")
		fmt.Scan(&dish)

		fmt.Printf("Enter %v's price: \n", dish)
		fmt.Scan(&dishPrice)

		//now populate the map
		ourMap[dish] = dishPrice

	} else {
		//else... if map Isn't empty; add too
		//prompt user to enter key and values repectively
		fmt.Println("Enter Dish: ")
		fmt.Scan(&dish)

		fmt.Printf("Enter %v's price: \n", dish)
		fmt.Scan(&dishPrice)

		//now populate the map
		ourMap[dish] = dishPrice
	}
}
