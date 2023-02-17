package main

import (
	"flag"
	"fmt"
	"math"

	conv "github.com/Shekina22/funtemp/conv"
)

// Definerer flag-variablene i hoved-"scope"
var (
	fahr float64
	out  string
	//funfacts   string
	Kelvin     float64
	Celsius    float64
	farhenheit float64
)

func init() {

	/*
	   Her er eksempler på hvordan man implementerer parsing av flagg.
	   For eksempel, kommando
	       funtemps -F 0 -out C
	   skal returnere output: 0°F er -17.78°C
	*/

	// Definerer og initialiserer flagg-variablene
	flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader fahrenheit")
	// Du må selv definere flag-variablene for "C" og "K"
	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")
	//flag.StringVar(&funfacts, "funfacts", "sun", "\"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")
	flag.Float64Var(&Celsius, "C", 0.0, "temperatur i grader celsius")
	flag.Float64Var(&Kelvin, "K", 0.0, "temperatur i grader Kelvin")
}

func main() {

	flag.Parse()

	// Her er noen eksempler du kan bruke i den manuelle testingen
	//fmt.Println(fahr, out, funfacts)

	//fmt.Println(isFlagPassed("out"))

	// Eksempel på enkel logikk
	/*if out == "C" && isFlagPassed("F") {
		// Kalle opp funksjonen FahrenheitToCelsius(fahr), som da
		// skal returnere °C
		fmt.Println("0°F er -17.78°C")
	}*/

	if out == "F" && isFlagPassed("K") {
		fmt.Println(Kelvin, "K er like", math.Round(conv.KelvinToFahrenheit(Kelvin)*100)/100, "F")
	}

	if out == "K" && isFlagPassed("F") {
		fmt.Println(Kelvin, "K er like", math.Round(conv.FahrenheitToKelvin(farhenheit)*100)/100, "K")
	}

	if out == "C" && isFlagPassed("K") {
		fmt.Println(Celsius, "Celsius er like", math.Round(conv.CelsiusToKelvin(Celsius)*100)/100, "C")
	}

	if out == "K" && isFlagPassed("C") {
		fmt.Println(Kelvin, "K er like", math.Round(conv.KelvinToCelsius(Kelvin)*100)/100, "C")
	}

	if out == "F" && isFlagPassed("C") {

		fmt.Println(Celsius, "C er like", math.Round(conv.CelsiusToFahrenheit(Celsius)*100)/100, "F")
	}

}

// Funksjonen sjekker om flagget er spesifisert på kommandolinje
// Du trenger ikke å bruke den, men den kan hjelpe med logikken
func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
