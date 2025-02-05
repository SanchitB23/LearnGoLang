package prices

import (
	"fmt"
	"max-tuts/price-calculator/conversion"
	"max-tuts/price-calculator/fileManager"
)

// CalculateTaxIncludedPrices calculates tax included prices.

type TaxIncludedPriceJob struct {
	IoManager         fileManager.FileManager `json:"-"`
	TaxRate           float64                 `json:"tax_rate"`
	InputPrices       []float64               `json:"input_prices,omitempty"`
	TaxIncludedPrices map[string]string       `json:"tax_included_prices,omitempty"`
}

func (job *TaxIncludedPriceJob) LoadData() {

	fileData, _ := job.IoManager.ReadFile()
	prices := conversion.StringsToFloats(fileData)
	job.InputPrices = prices
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()
	result := make(map[string]string)
	for _, price := range job.InputPrices {
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", price*(1+job.TaxRate))
	}
	job.TaxIncludedPrices = result
	err := job.IoManager.WriteJSON(job)
	if err != nil {
		return
	}
}

func NewTaxIncludedPricesJob(taxRate float64, fm fileManager.FileManager) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IoManager:   fm,
		TaxRate:     taxRate,
		InputPrices: []float64{},
	}
}
