package main

import "fmt"

//Declaring the license plates
var Licenseplates [3]string

//Give input
func main() {
	Licenseplates[0] = "32-H6T-1"
	Licenseplates[1] = "45-VET-4"
	Licenseplates[2] = "09-RQL-9"
	CheckLicenseplates()
}

//Check input
func CheckLicenseplates() {
	//Get user input
	var input string
	fmt.Println("Hoi en welkom! Vul hier je kenteken in:")
	fmt.Scan(&input)

	fmt.Println("Check het systeem met kenteken: " + input + "......")

	GiveInputToUser(input)
}

//Give back result
func GiveInputToUser(_input string) {
	var result bool = false

	for i := 0; i < len(Licenseplates); i++ {
		// checking if the array contains the given value
		if Licenseplates[i] == _input {
			// changing the boolean variable
			result = true
			break
		}
	}

	// printing the final result
	if result {
		fmt.Println("Array", Licenseplates, "Contains the given Value", _input)
	} else {
		fmt.Println("U heeft helaas geen toegang tot het parkeerterrein met kenteken: ", _input)
	}
}
