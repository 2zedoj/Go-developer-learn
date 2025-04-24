package main

import (
	"fmt"
	"convert/scale"
	"strconv"
	"bufio"
	"os"
    "strings"
)


func main () {
	var value float64
	var typeScaleCurrent, typeScaleDesired int
	reader := bufio.NewScanner(os.Stdin)

	// Pega o valor do usuário
	fmt.Print("Digite o valor a ser convertido: ")
	reader.Scan()
	value, err := strconv.ParseFloat(strings.TrimSpace(reader.Text()), 64) 
	if err != nil {
		fmt.Println("Error converting string to float:", err)
		return
	}

	// Qual a escala desse valor?
	fmt.Println("Qual a escala desse valor?")
	fmt.Println("   1 - Celsius    ")
	fmt.Println("   2 - Fahrenheit ")
	fmt.Println("   3 - Kelvin     ")
	reader.Scan()
	typeScaleCurrent, err = strconv.Atoi(strings.TrimSpace(reader.Text())) 
	if err != nil {
		fmt.Println("Error converting string to int:", err)
		return
	}
	// Valida se a opção existe
	if(typeScaleCurrent > 3 || typeScaleCurrent < 1){
		fmt.Println("Opção selecionada inválida de: " + strconv.Itoa(typeScaleCurrent))
		return
	}

	// Qual a escala que deseja converter
	fmt.Println("Qual a escala desejada?")
	fmt.Println("   1 - Celsius    ")
	fmt.Println("   2 - Fahrenheit ")
	fmt.Println("   3 - Kelvin     ")
	reader.Scan()
	typeScaleDesired, err = strconv.Atoi(strings.TrimSpace(reader.Text())) 
	if err != nil {
		fmt.Println("Error converting string to int:", err)
		return
	}
	if(typeScaleDesired > 3 || typeScaleDesired < 1){
		fmt.Println("Opção selecionada inválida de: " + strconv.Itoa(typeScaleDesired))
		return
	}

	// Cria o objeto
	scaleConvert := scale.Convert{
		Value: value,
		TypeScaleCurrent: scale.TypeScaleOptions(typeScaleCurrent),
		TypeScaleDesired: scale.TypeScaleOptions(typeScaleDesired),
	}

	// Pega o objeto do tipo atual
	scaleCurrent := scale.GetScale(scaleConvert.TypeScaleCurrent) 
	if scaleCurrent == nil {
        panic("Escala atual inválida!")
    }

	// Realiza a conversão
	switch scaleConvert.TypeScaleDesired {

	case scale.ScaleCelsius:{
		value = scaleCurrent.ConvertCelsius(&scaleConvert.Value)
	}

	case scale.ScaleFahrenheit:{
		value = scaleCurrent.ConvertFahrenheit(&scaleConvert.Value)
	}

	case scale.ScaleKelvin:{
		value = scaleCurrent.ConvertKelvin(&scaleConvert.Value)
	}
	
	default:
		panic("Opção selecionada inválida de: " + strconv.Itoa(typeScaleDesired))
	}

	fmt.Printf("Conversão de %s para %s é %.2f", scaleConvert.TypeScaleCurrent.String(), scaleConvert.TypeScaleDesired.String(), value)
}