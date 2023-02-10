package main

import (
	"fmt"
	"strconv"
	"time"
)

// Declaring the license plates
var Licenseplates [3]string

// Give input
func main() {
	Licenseplates[0] = "32-H6T-1"
	Licenseplates[1] = "45-VET-4"
	Licenseplates[2] = "09-RQL-9"
	CheckLicenseplates()
}

// Check input
func CheckLicenseplates() {
	//Get user input
	var input string
	fmt.Println("Hoi en welkom! Vul hier je kenteken in:")
	fmt.Scan(&input)

	fmt.Println("Check het systeem met kenteken: " + input + "......")

	GiveInputToUser(input)
}

// Give back result
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
		DisplayMessage()
	} else {
		fmt.Println("U heeft helaas geen toegang tot het parkeerterrein met kenteken: ", _input)
	}
}

func DisplayMessage() {
	var groet string

	//getting the time. Total, hours and minutes
	currentTime := time.Now()
	hour := currentTime.Hour()
	minutes := currentTime.Minute()

	//adds 0 when the time gets back as single number
	minuteString := fmt.Sprintf("%02d", minutes)

	//Displaying the time to the user
	fmt.Println("Current time: ", hour, ":", minuteString)

	//Converting the time to a readable/useable string
	convertedTime := strconv.Itoa(hour) + minuteString

	//Convert string to int
	convertedTimeAsInt, err := strconv.Atoi(convertedTime)
	if err != nil {
		// ... handle error
		panic(err)
	}

	//Printing the converted time [DEBUG ONLY]
	fmt.Println(convertedTimeAsInt)

	//Setting greeting based on time of day.
	var defaultText string = "Welkom bij vakantie fonteyn!"
	var nightText string = "Sorry, de parkeerplaats is ’s nachts gesloten"
	var textUsed string

	if convertedTimeAsInt >= 700 && convertedTimeAsInt <= 1159 {
		groet = "Goedemorgen! "
		textUsed = defaultText
	}
	if convertedTimeAsInt >= 1200 && convertedTimeAsInt <= 1759 {
		groet = "Goedemiddag! "
		textUsed = defaultText
	}
	if convertedTimeAsInt >= 1800 && convertedTimeAsInt <= 2259 {
		groet = "Goedeavond! "
		textUsed = defaultText
	}
	if convertedTimeAsInt >= 2300 && convertedTimeAsInt <= 659 {
		groet = ""
		textUsed = nightText
	}
	//Displaying the greeting
	fmt.Println(groet + textUsed)
}
