package main

import (
	"fmt"

	".cli-demos/price-calculator-cli/prices"
)

func main() {

	taxRate := []float64{0.05, 0.1, 0.15, 0.2}
	result := make(map[float64][]float64)
	for _, tax := range taxRate {
		taxJob := prices.NewTaxIncludedPrice(tax)
		taxJob.Process()
	}

	fmt.Println(result)
}
