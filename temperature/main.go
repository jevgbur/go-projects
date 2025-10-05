package main

import (
	"fmt"
	"os"
	"strconv"
)

type Temperature struct {
	Value float64
	Unit  string
}

func celsiusToFahrenheit(t *Temperature) *Temperature {
	t.Value = (t.Value * 1.8) + 32
	t.Unit = "F"
	return &Temperature{Value: t.Value, Unit: t.Unit}
}

func fahrenheitToCelsius(t *Temperature) *Temperature {
	t.Value = (t.Value - 32) / 1.8
	t.Unit = "C"
	return &Temperature{Value: t.Value, Unit: t.Unit}
}

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
	temp := &Temperature{Value: value, Unit: unit}

	var convertedTemp *Temperature

	switch temp.Unit {
	case "C":
		convertedTemp = celsiusToFahrenheit(temp)
	case "F":
		convertedTemp = fahrenheitToCelsius(temp)
	default:
		fmt.Println("Invalid unit. Please provide 'C' for Celsius or 'F' for Fahrenheit.")
		os.Exit(1)
	}
	fmt.Printf("%.0f %s\n", convertedTemp.Value, convertedTemp.Unit)

}
