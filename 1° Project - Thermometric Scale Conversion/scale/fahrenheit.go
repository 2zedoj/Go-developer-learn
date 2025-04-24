package scale

type Fahrenheit struct{}

func(f Fahrenheit) ConvertCelsius(value *float64) float64 {
	return (*value-32)*5/9
}

func(f Fahrenheit) ConvertFahrenheit(value *float64) float64 {
	return *value
}

func(f Fahrenheit) ConvertKelvin(value *float64) float64 {
	return (*value - 32)*5/9 + 273.15
}