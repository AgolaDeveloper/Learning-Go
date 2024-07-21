// this program file deletes items from the map
package main

import "fmt"

//func deletes item from the map
//takes 2 parameters, ourMap and isMapEmpty
func deleteItem(ourMap map[string]int, isMapEmpty bool) {

	var item2Delete string
	//can't delete an item if map is empty
	//thus, checking first whether map is empty

	if isMapEmpty {
		//if map is empty... print->
		fmt.Printf("\nOoops...Can't Perform any Deletion Operation \nMap is empty\n")
	} else {
		//else... if map isn't empty go ahead and delete the item
		fmt.Println("Item to Delete?: ")
		fmt.Scan(&item2Delete)

		delete(ourMap, item2Delete)
	}
}
