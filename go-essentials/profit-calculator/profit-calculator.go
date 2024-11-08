package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	// var revenue, expenses, taxRate float64

	// fmt.Print("Revenue: ")
	// fmt.Scan(&revenue)

	// fmt.Print("Expenses: ")
	// fmt.Scan(&expenses)

	// fmt.Print("Tax Rate: ")
	// fmt.Scan(&taxRate)

	revenue, err1 := inputText("Revenue: ")



	expenses, err2 := inputText("Expenses: ")
	taxRate, err3 := inputText("Tax Rate: ")

	if err1 != nil || err2 != nil || err3 != nil{
		fmt.Println("ERROR")
		fmt.Println(err1)
		fmt.Println("------")
		panic("Can't continue. Sorry.")
	}

	// ebt := revenue - expenses
	// profit := ebt * (1 - taxRate / 100)
	// ratio := ebt / profit

	ebt, profit, ratio := calculateProfit(revenue, expenses, taxRate)

	// fmt.Println(ebt)
	// fmt.Println(profit)
	// fmt.Println(ratio)
	storeResults(ebt, profit, ratio)
}

func inputText(text string)(float64, error){
	fmt.Print(text)
	var input float64
	fmt.Scan(&input)

	if input <= 0 {
		return 0, errors.New("input should be positive")
	}

	return input, nil
}

func calculateProfit(revenue, expenses, taxRate float64)(ebt float64, profit float64, ratio float64){
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate / 100)
	ratio = ebt / profit
	return ebt, profit, ratio
}

func storeResults(ebt, profit, ratio float64){
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.1f\n", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}