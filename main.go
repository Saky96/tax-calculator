package main

import (
	"fmt"
	"github.com/Saky96/tax-calculator/prices"
	"github.com/Saky96/tax-calculator/utils"
)

//type taxPricesMap

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		fm := utils.NewFileManager("prices.txt", fmt.Sprintf("results/result_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(taxRate, fm)
		priceJob.Process()
	}
}
