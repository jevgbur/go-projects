package main

import (
	"fmt"
	"os"
	"strconv"
	"temp_converter/conversions"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <value> <unit (C/F)>")
		os.Exit(1)
	}

	strValue := os.Args[1]
	unit := os.Args[2]

	value, err := strconv.ParseFloat(strValue, 64)
	if err != nil {
		fmt.Println("Invalid temperature value. Please enter a numeric value.")
		os.Exit(1)
	}

	// Create the pointer to Temperature using parsed inputs
	temp := conversions.Temperature{Value: value, Unit: unit}

	var out *conversions.Temperature

	switch temp.Unit {
	case "C":
		out = conversions.CelsiusToFahrenheit(&temp)
	case "F":
		out = conversions.FahrenheitToCelsius(&temp)
	default:
		fmt.Println("Invalid unit. Please provide 'C' for Celsius or 'F' for Fahrenheit.")
		os.Exit(1)
	}
	fmt.Printf("%.0f %s\n", out.Value, out.Unit)
}
