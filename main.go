package main

import (
	"fmt"
	"github.com/Saky96/tax-calculator/prices"
	"github.com/Saky96/tax-calculator/utils"
)

//type taxPricesMap

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	doneChannels := make([]chan bool, len(taxRates)) // Create a channel for each tax rate to signal completion

	for i, taxRate := range taxRates {
		doneChannels[i] = make(chan bool) // Initialize the channel for each goroutine
		fm := utils.NewFileManager("prices.txt", fmt.Sprintf("results/result_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(taxRate, fm)

		go priceJob.Process(doneChannels[i]) // Start the goroutine for each tax rate
		//if err != nil {
		//	return
		//}
	}

	for _, doneChannel := range doneChannels {
		<-doneChannel // Wait for each goroutine to finish
	}
}
