package main

import (
	"GO/services"
	"fmt"
)

func main() {
	//init method instances

	arithmetic := services.ArithmeticService{}
	input := services.UserInput{}

	//user input
	num := input.ReadNumbers()
	oper := input.ReadOperand()

	//operations
	switch oper {
	case "+":
		sum := arithmetic.Sum(num[0], num[1])
		fmt.Println("Sum:", sum)

	case "-":
		diff := arithmetic.Diff(num[0], num[1])
		fmt.Println("Difference: ", diff)

	case "*":
		prod := arithmetic.Product(num[0], num[1])
		fmt.Println("Product: ", prod)

	case "/":
		div := arithmetic.Div(num[0], num[1])
		fmt.Println("Quotient: ", div)

	default:
		fmt.Println("Wrong operator input. Please use either of these (+,-,*,/)")
	}

}
