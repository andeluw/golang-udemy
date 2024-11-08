package main

import (
	// "example.com/price-calc/cmdmanager"
	"fmt"

	"example.com/price-calc/filemanager"
	"example.com/price-calc/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate * 100))
		// cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		// err := priceJob.Process()
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)
		go priceJob.Process(doneChans[index], errorChans[index])

		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }
	}

	for index := range taxRates {
		select {
			case err := <-errorChans[index]:
				if err != nil {
					fmt.Println(err)
				}
			case <-doneChans[index]:
				fmt.Println("Done!")
		}

	}

	// for _, doneChan := range doneChans {
	// 	<- doneChan
	// }
}