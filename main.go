package main

import (
	"fmt"
	"github.com/Saky96/tax-calculator/prices"
	"github.com/Saky96/tax-calculator/utils"
)

//type taxPricesMap

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	doneChannels := make([]chan bool, len(taxRates))   // Create a channel for each tax rate to signal completion
	errorChannels := make([]chan error, len(taxRates)) // Create a channel for each tax rate to signal errors

	for i, taxRate := range taxRates {
		doneChannels[i] = make(chan bool)   // Initialize the channel for each goroutine
		errorChannels[i] = make(chan error) // Initialize the error channel for each goroutine
		fm := utils.NewFileManager("prices.txt", fmt.Sprintf("results/result_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(taxRate, fm)

		//go priceJob.Process(doneChannels[i]) // Start the goroutine for each tax rate
		go priceJob.Process(doneChannels[i], errorChannels[i]) // Start the goroutine for each tax rate with error handling
		//if err != nil {
		//	return
		//}
	}

	//for _, doneChannel := range doneChannels {
	//	<-doneChannel // Wait for each goroutine to finish
	//}

	for i, _ := range taxRates {
		select {
		case val := <-doneChannels[i]: // Wait for the goroutine to finish
			fmt.Println("Done Channel: ", val)
		case err := <-errorChannels[i]: // Handle any errors that occurred in the goroutine
			fmt.Println("Error: ", err)
		}
	}
}
