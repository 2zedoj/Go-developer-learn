package main

import "fmt"


type intervalo struct{
	valorDe int8
	valorAte int8
	divisorPin int8
	divisorPan int8
} 

func jogoPinPan(in intervalo) {
	for i := in.valorDe; i < in.valorAte; i++ {
		if i % in.divisorPin == 0{
			fmt.Println("Pin ")
		}else if i % in.divisorPan == 0{
			fmt.Println("Pan ")
		}

		if (i % in.divisorPin == 0) && (i % in.divisorPan == 0) {
			fmt.Println("Pin Pan")
		}
	}
}

func main(){

	i := intervalo{valorDe: 1, valorAte: 100, divisorPin: 3, divisorPan: 5}

	jogoPinPan(i)
}