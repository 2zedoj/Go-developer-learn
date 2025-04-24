package scale

type Kelvin struct{}

func(k Kelvin) ConvertCelsius(value *float64) float64 {
	return *value - 273.1 
}

func(k Kelvin) ConvertFahrenheit(value *float64) float64 {
	return (*value-273.15)*9/5+32
}

func(k Kelvin) ConvertKelvin(value *float64) float64 {
	return *value
}