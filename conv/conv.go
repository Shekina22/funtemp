package conv

/*
  I denne pakken skal alle konverteringfunksjonene
  implementeres. Bruk engelsk.
    FarhenheitToCelsius
    CelsiusToFahrenheit
    KelvinToFarhenheit
    ...
*/
//konverterer Farhnrenheit til Celsius
func FahrenheitToCelsius(value float64) float64 {
	return (value - 32) * 5 / 9
}

//konverterer Celsius til Farhrenheit
func CelsiusToFahrenheit(value float64) float64 {
	return (value * 9 / 5) + 32
}

//konverterer til Kelvin til Fahrenheit
func KelvinToFahrenheit(value float64) float64 {
	return value*1.8 - 459.67
}

//konverterer Celsius til Kelvin
func CelsiusToKelvin(value float64) float64 {
	return value + 273.15
}

//konverterer Farhrenheit til Kelvin
func FahrenheitToKelvin(value float64) float64 {
	return (value-32)*5/9 + 273.15
}

//konverterer til Celsius
func KelvinToCelsius(value float64) float64 {
	return value - 273.15
}
