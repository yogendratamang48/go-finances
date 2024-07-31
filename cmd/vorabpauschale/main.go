package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Etf struct {
	Name       string  `json:"name"`
	StartValue float32 `json:"startValue"`
	EndValue   float32 `json:"endValue"`
}

// Vorabpauschale (Advance lump sum) is 70% of the base interest rate (2.55% for 2023) of the value at the start of the year
func (n Etf) calculateAdvanceLumpsum() float32 {
	return n.StartValue * 0.70 * 0.0255
}

func (n Etf) calculateUnrealizedGain() float32 {
	return n.EndValue - n.StartValue
}

func (n Etf) calculateTaxableAmount() float32 {
	if n.calculateAdvanceLumpsum() > n.calculateUnrealizedGain() {
		return n.calculateUnrealizedGain()
	}
	return n.calculateAdvanceLumpsum()
}

func (n Etf) calculateTax() float32 {

	return n.calculateTaxableAmount() * 26.375 / 100
}

// load etfs array from json file
func loadEtfs() []Etf {
	var etfs []Etf
	file, err := os.ReadFile("./cmd/vorabpauschale/etfs.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(file, &etfs)
	if err != nil {
		log.Fatal(err)
	}
	return etfs
}

func main() {
	etfs := loadEtfs()
	var totalTax float32 = 0.0
	for _, etf := range etfs {
		fmt.Println("Accumulating Stock: ", etf.Name)
		fmt.Println("Value at the start of year: ", etf.StartValue)
		fmt.Println("Value at the end of year: ", etf.EndValue)
		fmt.Println("Unrealized gain: ", etf.calculateUnrealizedGain())
		fmt.Println("Vorabpauschale(Advance lump-sum): ", etf.calculateAdvanceLumpsum())
		fmt.Println("Taxable amount: ", etf.calculateTaxableAmount())
		fmt.Println("Tax amount: ", etf.calculateTax())
		totalTax += etf.calculateTax()
		fmt.Println("------------------------------------------------")
	}
	fmt.Println("Total Tax: ", totalTax)
}
