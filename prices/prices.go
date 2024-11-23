package prices

import (
	"example.com/practice_app/file_handling"
	"fmt"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64						`json:"tax_rate"`
	InputPrices       []float64						`json:"input_prices"`
	TaxIncludedPrices map[string]string				`json:"tax_included_prices"`
	FileHandler		  file_handling.FileHandler		`json:"-"`
}

func NewJob(taxRate float64, fileHandler file_handling.FileHandler) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate: taxRate,
		FileHandler: fileHandler,
	}
}

func (job *TaxIncludedPriceJob) Process() {
	job.InputPrices = job.FileHandler.LoadData()
	result := make(map[string]string)

	for _, price := range job.InputPrices {
		calculation := price * (1 + job.TaxRate) 
		TaxedPrice := fmt.Sprintf("%.2f", calculation)
		result[fmt.Sprintf("%v", price)] = TaxedPrice
	}

	job.TaxIncludedPrices = result

	fileName := fmt.Sprintf("result_%v.json", job.TaxRate * 100)

	err := job.FileHandler.WriteJson(job, fileName)
	if err != nil {
		fmt.Printf("the JSON thingy failed: %v\n", err)
	}



}

