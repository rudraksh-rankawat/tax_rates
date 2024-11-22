package main

import (
	"example.com/practice_app/prices"
)

func main() {

	taxRates := []float64{0.01, 0.07, 0.15}

	job1 := prices.NewJob(taxRates[0])

	job1.Process()

}