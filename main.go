package main

import "fmt"

//type taxPricesMap

func main() {
	prices := []float32{10, 20, 30}
	taxRates := []float32{0, 0.07, 0.1, 0.15}

	taxRatePricesMap := make(map[float32][]float32)

	for _, taxRate := range taxRates {
		taxPrice := make([]float32, len(prices))
		for j, price := range prices {
			taxPrice[j] = price * (1 + taxRate)
		}
		taxRatePricesMap[taxRate] = taxPrice
	}

	fmt.Println(taxRatePricesMap)
}
