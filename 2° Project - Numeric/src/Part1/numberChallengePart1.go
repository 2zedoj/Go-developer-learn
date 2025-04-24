package main

import "fmt"


type intervalo struct{
	valorDe int8
	valorAte int8
	divisor int8
} 

func divisivelDe(in intervalo) {
	for i := in.valorDe; i < in.valorAte; i++ {
		if i % in.divisor == 0 {
			fmt.Printf("Valor %d é divisor de %d \n", i, in.divisor)
		}
	}
}

func main(){

	i := intervalo{valorDe: 1, valorAte: 100, divisor: 3}

	divisivelDe(i)
}