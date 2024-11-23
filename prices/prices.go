package prices

import (
	"bufio"
	"fmt"
	"os"
	"example.com/practice_app/conversion"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func NewJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate: taxRate,
	}
}

func (job TaxIncludedPriceJob) Process() {
	job.LoadData("price_data.txt")
	result := make(map[string]float64)

	for priceIndex, price := range job.InputPrices {
		TaxedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%v", priceIndex+1)] = TaxedPrice
	}

	fmt.Println(result)
}

func (job *TaxIncludedPriceJob) LoadData(fileName string) {
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Cannot open the file")
		fmt.Println(err)
		return
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	price_data_str := []string{}

	for scanner.Scan() {
		price_data_str = append(price_data_str, scanner.Text())
	}
	file.Close()
	
	job.InputPrices = conversion.ConvertStrToFloatMap(price_data_str)
}
