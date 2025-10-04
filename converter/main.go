package main

import "fmt"

type distance float64

func main() {
	var miles distance
	miles = 30

	km := miles.ToKm()
	fmt.Println("Kilometers:", km)
	fmt.Println("Miles:", miles)
}

func (d distance) ToKm() distance {
	return d * 1.6
}
