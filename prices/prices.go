package prices

import (
	"fmt"
	"github.com/Saky96/tax-calculator/utils"
)

type MapOfStringAndString map[string]string

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices MapOfStringAndString
	IOManager         utils.FileManager
}

func (job *TaxIncludedPriceJob) LoadData() {

	lines, err := job.IOManager.ReadFile()
	prices, err := utils.StingsToFloats(lines)

	if err != nil {
		fmt.Println(err)
		return
	}

	job.InputPrices = prices
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()
	result := make(MapOfStringAndString)
	for _, price := range job.InputPrices {
		taxCalculatedPrices := fmt.Sprintf("%.2f", price*(1+job.TaxRate))
		result[fmt.Sprintf("%.2f", price)] = taxCalculatedPrices
	}

	fmt.Println(result)

	job.TaxIncludedPrices = result

	//err := job.IOManager.WriteFile(job, fmt.Sprintf("results/result_%.0f.json", job.TaxRate*100))
	err := job.IOManager.WriteFile(job)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func NewTaxIncludedPriceJob(taxRate float64, fm utils.FileManager) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
		IOManager:   fm,
	}
}
