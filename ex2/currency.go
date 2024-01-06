package main

import "fmt"

const comision = 0.01

type Currency struct {
	name       string
	valueToUSD float32
	min        float32
	max        float32
}

var currencys = map[string]Currency{
	"CLP": {name: "CLP", valueToUSD: 889.95, min: 10, max: 100},
	"ARS": {name: "ARS", valueToUSD: 811.20, min: 10, max: 100},
	"USD": {name: "USD", valueToUSD: 1, min: 10, max: 100},
	"EUR": {name: "EUR", valueToUSD: 0.91, min: 10, max: 100},
	"TRY": {name: "TRY", valueToUSD: 29.82, min: 10, max: 100},
	"GBP": {name: "GBP", valueToUSD: 0.79, min: 10, max: 100},
}

func main() {
	welcome()

	for {

		currency1 := getInput("Escoja la divisa inicial: ")
		currency2 := getInput("Escoja la divisa a convertir: ")
		amount := getAmout("Escoja la cantidad a convertir: ")
		to := convert(currency1, currency2, amount)
		fmt.Println("\n", amount, " ", currency1, " son ", to, " ", currency2)
		if getInput("Desea Retirar el valor ([S]Si - [N]No ): ") == "S" {
			var amoutToWithDraw float32 = to - (to * comision)
			fmt.Println("Retirastes ", amoutToWithDraw, currency2)
		}

		if getInput("Desea continuar con otra transacion ([S]Si - [N]No ): ") == "S" {
			continue
		}
		break
	}
	fmt.Print("\nADIOS\n\n")
}

func welcome() {
	fmt.Print("\nWelcome to currency exchange\n\n")
}

func getInput(message string) string {
	fmt.Print(message)
	var input string
	fmt.Scan(&input)
	return input
}

func getAmout(message string) float32 {
	fmt.Print(message)
	var amount float32
	fmt.Scan(&amount)
	return amount
}

func convert(currency1 string, currency2 string, amount float32) float32 {
	exc1, existe1 := currencys[currency1]
	exc2, existe2 := currencys[currency2]
	if existe1 && existe2 {
		return amount / exc1.valueToUSD * exc2.valueToUSD
	}
	return 0
}
