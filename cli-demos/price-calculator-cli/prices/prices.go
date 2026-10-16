package prices

import (
	"bufio"
	"fmt"
	"os"
)

type TaxIncludedPrice struct { // Struct to hold tax-included price information
	TaxRate           float64
	Prices            []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPrice) LoadData() {
	// Method to load data from a file
	file, err := os.Open("prices.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		fmt.Println("Error reading file:", err)
		file.Close()
		return
	}

}

func (job *TaxIncludedPrice) Process() {
	// Method to calculate tax-included prices
	result := make(map[string]float64)
	for _, price := range job.Prices {
		result[fmt.Sprintf("%.2f", price)] = price * (1 + job.TaxRate)
	}

	fmt.Println(result)
}

func NewTaxIncludedPrice(taxRate float64) *TaxIncludedPrice {
	return &TaxIncludedPrice{
		TaxRate: taxRate,
		Prices:  []float64{10, 20, 30, 40},
	}
}
