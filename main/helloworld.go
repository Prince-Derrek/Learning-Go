package main

import (
	"GO/services"
	"fmt"
)

func main() {
	//user input
	numberService := services.UserInput{}
	num := numberService.Variant()

	operandService := services.UserInput{}
	oper := operandService.Operand()

	//operations
	switch oper {
	case "+":
		sumService := services.ArithmeticService{}
		sum := sumService.Sum(num[0], num[1])
		fmt.Println("Sum:", sum)

	case "-":
		diffService := services.ArithmeticService{}
		diff := diffService.Diff(num[0], num[1])
		fmt.Println("Difference: ", diff)

	case "*":
		prodService := services.ArithmeticService{}
		prod := prodService.Product(num[0], num[1])
		fmt.Println("Product: ", prod)

	case "/":
		divService := services.ArithmeticService{}
		div := divService.Div(num[0], num[1])
		fmt.Println("Quotient: ", div)

	default:
		fmt.Println("Wrong operator input. Please use either of these (+,-,*,/)")
	}

}
