// Using the package main
package main

import (
	"fmt"
	"strconv"
	"time"
)

// Function called main that prints out a line.
func main() {
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
