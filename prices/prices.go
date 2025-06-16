package prices

import (
	"fmt"
	"github.com/Saky96/tax-calculator/utils"
)

type MapOfStringAndString map[string]string

type TaxIncludedPriceJob struct {
	TaxRate           float64              `json:"tax_rate"`
	InputPrices       []float64            `json:"input_prices"`
	TaxIncludedPrices MapOfStringAndString `json:"taxed_prices"`
	IOManager         utils.FileManager    `json:"-"`
}

func (job *TaxIncludedPriceJob) LoadData() error {

	lines, err := job.IOManager.ReadFile()
	prices, err := utils.StingsToFloats(lines)

	if err != nil {
		fmt.Println(err)
		return err
	}

	job.InputPrices = prices

	return nil
}

func (job *TaxIncludedPriceJob) Process(doneChannel chan bool, errChannel chan error) {
	err := job.LoadData()

	//errChannel <- errors.New("Adding errors for testing error channels") // Simulating an error for testing purposes

	if err != nil {
		errChannel <- err
		return
	}
	result := make(MapOfStringAndString)
	for _, price := range job.InputPrices {
		taxCalculatedPrices := fmt.Sprintf("%.2f", price*(1+job.TaxRate))
		result[fmt.Sprintf("%.2f", price)] = taxCalculatedPrices
	}

	fmt.Println(result)

	job.TaxIncludedPrices = result

	//err := job.IOManager.WriteFile(job, fmt.Sprintf("results/result_%.0f.json", job.TaxRate*100))
	job.IOManager.WriteFile(job)
	doneChannel <- true // Signal that the job is done
}

func NewTaxIncludedPriceJob(taxRate float64, fm utils.FileManager) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
		IOManager:   fm,
	}
}
