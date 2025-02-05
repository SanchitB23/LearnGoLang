package main

import (
	"fmt"
	"max-tuts/price-calculator/fileManager"
	"max-tuts/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.1, 0.2, 0.3, 0.4, 0.5}

	for _, taxRate := range taxRates {
		prices.NewTaxIncludedPricesJob(taxRate, fileManager.NewFileManager("prices.txt", fmt.Sprintf("tax_included_prices_%d.json", int(taxRate*100)))).Process()
	}
}
