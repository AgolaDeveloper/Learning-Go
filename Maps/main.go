package main

// this program file is the entry point to our Maps' executable programs

func main() {
	//first initialize ourMap into a local map variable
	ourStruct := OurMap()

	//add item
	AddItem(ourStruct, IsMapEmpty)

	//delete an item
	DeleteItem(ourStruct, IsMapEmpty)

}
