package main

import "fmt"

func main() {
	CalculatorInstance.inputValue(3)
	CalculatorInstance.chooseActions("+")
	CalculatorInstance.inputValue(5)
	CalculatorInstance.chooseActions("*")
	CalculatorInstance.inputValue(2)
	fmt.Println(CalculatorInstance.getResult())
}
