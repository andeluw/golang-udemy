package main

import "fmt"

type customString string

func (text customString) log() {
	fmt.Print(text)
}

func main() {
	var name customString

	name = "Max"
	name.log()
}