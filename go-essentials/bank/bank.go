package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Call us 24/7", randomdata.PhoneNumber())

	loop:
	for {
		balance, err := fileops.GetFloatFromFile(accountBalanceFile, 1000)

		if err != nil {
			fmt.Println("ERROR")
			fmt.Println(err)
			fmt.Println("----------")
			panic("Can't continue, sorry.")
		}
		
		presentOptions();

		var choice int
		fmt.Scan(&choice)

		switch choice {
			case 1:
				fmt.Printf("Your balance is %v\n", balance)
			case 2:
				depositMoney(balance)
			case 3:
				withdrawMoney(balance)
			case 4:
				break loop
			default:
				fmt.Println("Choice not valid!")		
		}
	}

	fmt.Println("Thanks for choosing our bank.")
}

func depositMoney(balance float64){
	var depositAmount float64
	fmt.Print("Deposit amount: ")
	fmt.Scan(&depositAmount)
	if(depositAmount <= 0){
		fmt.Println("Deposit amount should be positive!")
		return
	}
	balance += depositAmount
	fmt.Printf("Balance updated. New amount: %v\n", balance)
	fileops.WriteFloatToFile(balance, accountBalanceFile)
}

func withdrawMoney(balance float64){
	var withdrawalAmount float64
	fmt.Print("Withdrawal amount: ")
	fmt.Scan(&withdrawalAmount)
	if(withdrawalAmount <= 0){
		fmt.Println("Withdrawal amount should be positive!")
		return
	} else if(withdrawalAmount > balance){
		fmt.Println("Withdrawal amount should be less than balance!")
	}
	balance -= withdrawalAmount
	fmt.Printf("Balance updated. New amount: %v\n", balance)
	fileops.WriteFloatToFile(balance, accountBalanceFile)
}