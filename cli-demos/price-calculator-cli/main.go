package main

import (
	"fmt"
)

func main() {
	prices := []float64{10, 20, 30, 40}

	taxRate := []float64{0.05, 0.1, 0.15, 0.2}
	result := make(map[float64][]float64)
	for _, tax := range taxRate {
		taxIncludedPrices := make([]float64, len(prices))
		for j, price := range prices {
			taxIncludedPrices[j] = price * (1 + tax)
		}
		result[tax] = taxIncludedPrices
	}

	fmt.Println(result)
}
