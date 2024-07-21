// this program file checks if our Map is Empty
package main

func IsMapEmpty(ourMap map[string]int) bool {

	if len(ourMap) == 0 {
		//if its len is zero; YES... ourMap is empty
		return true
	} else {
		//else if its len aint ZERO; NO... ourMap isn't empty
		return false
	}
}
