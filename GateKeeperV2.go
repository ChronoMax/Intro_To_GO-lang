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

	//Time
	currentTime := time.Now()
	hour := currentTime.Hour()
	minutes := currentTime.Minute()

	//Displaying the time to the user
	fmt.Println("Current time: ", hour, ":", minutes)

	//Converting the time to a readable/useable string
	convertedTime := strconv.Itoa(hour) + strconv.Itoa(minutes)

	//Printing the converted time [DEBUG ONLY]
	//fmt.Println(convertedTime)

	//Setting greeting based on time of day.
	if convertedTime >= "700" && convertedTime <= "1159" {
		groet = "Goedemorgen! "
	}
	if convertedTime >= "1200" && convertedTime <= "1759" {
		groet = "Goedemiddag! "
	}
	if convertedTime >= "1800" && convertedTime <= "2300" {
		groet = "Goedeavond! "
	}
	//Displaying the greeting + default string to the user
	fmt.Println(groet + "Welkom bij vakantie fonteyn!")
}
