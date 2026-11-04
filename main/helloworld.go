package main

import (
	"GO/services"
	"fmt"
)

func Calculate() string {
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

	continuance := ""
	fmt.Println("\n\nWould you like to proceed (y if Yes, n if No)")
	fmt.Scan(&continuance)

	return continuance
}

func main() {
	fmt.Println("Welcome to my new calculator! \n")

	continuance := "y"

	for continuance == "y" {
		continuance = Calculate()
	}

	fmt.Println("\n\nThank you for using My Calculator")

}
