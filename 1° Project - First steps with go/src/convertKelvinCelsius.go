package main

import "fmt"

var K = 373.2

func main () {
	tempC := K - 273

	fmt.Printf("Ponto de Ebulição em Kelvin: %g°K e em Celsius: %.2f°C", K, tempC)
}