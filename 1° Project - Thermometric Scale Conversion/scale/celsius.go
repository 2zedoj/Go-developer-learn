package scale

type Celsius struct{}

func(c Celsius) ConvertCelsius(value *float64) float64 {
	return *value*9/5 + 32
}

func(c Celsius) ConvertFahrenheit(value *float64) float64 {
	return *value*9/5 + 32
}

func(c Celsius) ConvertKelvin(value *float64) float64 {
	return *value
}
