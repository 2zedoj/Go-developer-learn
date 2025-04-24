package scale

type TypeScaleOptions int

const (
	ScaleCelsius TypeScaleOptions = 1
	ScaleFahrenheit TypeScaleOptions = 2
	ScaleKelvin TypeScaleOptions = 3
)

type Convert struct{
	Value float64
	TypeScaleCurrent TypeScaleOptions
	TypeScaleDesired TypeScaleOptions
}

type ScaleConvert interface 
{
	ConvertCelsius(value *float64) float64
	ConvertFahrenheit(value *float64) float64
	ConvertKelvin(value *float64) float64
}

func GetScale(currentScale TypeScaleOptions) ScaleConvert {
    switch currentScale {
    case ScaleCelsius:
        return Celsius{}
    case ScaleFahrenheit:
        return Fahrenheit{}
    case ScaleKelvin:
        return Kelvin{}
    default:
        return nil
    }
}

func (s TypeScaleOptions) String() string {
	switch s {
	case ScaleCelsius:
		return "Celsius"
	case ScaleFahrenheit:
		return "Fahrenheit"
	case ScaleKelvin:
		return "Kelvin"
	default:
		return "Desconhecido"
	}
}