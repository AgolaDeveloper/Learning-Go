package main

//here WE HELP USERS WITH TEMP-CONVERSION
//USER MUST ENTER THE TEMP AND TARGET TEMP
import "fmt"

//functions that help with the conversions
//...both take two parameters from user;
//userTemp and Target temp...plus
//a third constant parameter

//this helps with converting Celsius->Fah

func conCelsius(userT, k float64) {
	finalT := userT * k

	fmt.Printf("%v C -> Fah: %v \n", userT, finalT)
}

//this helps with converting Fah->Celsius
func conFah(userTemp, kon float64) {
	finalTemp := userTemp / kon

	fmt.Printf("%v Fah -> C: %v \n", userTemp, finalTemp)
}

func main() {

	//formular for Celcius->Fah Conversion
	//... Fah=celcius*121.63

	//formular for Fah->Celcius Conversion
	//... Cel=Fah/121.63

	const k = 121.63

	var yourTemp float64
	var conversionFit int

	//user will do the following:
	//1. Enter the Temp they wanna convert [yourTemp]
	//2. Enter the target Temp they need their original temp converted to

	//3. finally trigger display of the conversion/result

	//FIRST MAKING IT EASY FOR USER TO NAVIGATE U/I

convert:
	fmt.Printf("\nCHOOSE THE BEST FIT FOT YOUR CONVERSION\n ************************\n")
	fmt.Println("1. Celcius to Fah")
	fmt.Println("2. Fah to Celcius")

	fmt.Scan(&conversionFit)

	//then we use switch
	switch conversionFit {
	case 1:
		fmt.Println("ENTER Temp You wanna Convert (C): ")
		fmt.Scan(&yourTemp)

		conCelsius(yourTemp, k)

	case 2:
		fmt.Println("ENTER Temp You wanna Convert (Fah): ")
		fmt.Scan(&yourTemp)

		conFah(yourTemp, k)

	default:
		fmt.Println("INVALID CHOICE: Retry....")
		goto convert
	}

}
