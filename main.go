package main

import (
	"example.com/practice_app/prices"
	"example.com/practice_app/file_handling"
)

func main() {

	taxRates := []float64{0.01, 0.07, 0.15, 0.90, 1.1}



	for _, taxRate := range taxRates {
		FH := file_handling.NewFileHandler("price_data.txt")
		job := prices.NewJob(taxRate, *FH)
		job.Process()
	}
}