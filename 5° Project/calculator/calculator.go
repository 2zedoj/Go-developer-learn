package calculator

import (
	//"fmt"
)

func sum(i ...float64) float64{
	total := 0.0
	
	for _, r := range i{
		total += r
	}

	return total
} 

func subtract(i ...float64) float64{
	total := 0.0
	
	for _, r := range i{
		total -= r
	}

	return total
} 

func division(divisor float64, dividendo float64) float64{
	if dividendo == 0 {
		panic("divisão por zero")
	}

	return divisor/dividendo
} 

func multiplication(i ...float64) float64{
	total := 1.0
	
	for _, r := range i{
		total *= r
	}

	return total
} 
