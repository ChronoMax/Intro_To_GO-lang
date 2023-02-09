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

	currentTime := time.Now()
	hour := currentTime.Hour()
	minutes := currentTime.Minute()

	fmt.Println("Current time: ", hour, ":", minutes)

	convertedTime := strconv.Itoa(hour) + strconv.Itoa(minutes)

	fmt.Println(convertedTime)

	if convertedTime >= "700" && convertedTime <= "1159" {
		groet = "Goedemorgen! "
	}
	if convertedTime >= "1200" && convertedTime <= "1759" {
		groet = "Goedemiddag! "
	}
	if convertedTime >= "1800" && convertedTime <= "2300" {
		groet = "Goedeavond! "
	}
	fmt.Println(groet + "Welkom bij vakantie fonteyn!")
}
