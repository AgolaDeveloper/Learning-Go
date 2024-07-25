package main

// this program file is the entry point to our Maps' executable programs

func main() {
	//first initialize ourMap into a local map variable
	ourStruct := ourMap()

	//add item
	addItem(ourStruct, IsMapEmpty)

	//delete an item
	deleteItem(ourStruct, IsMapEmpty)

}
