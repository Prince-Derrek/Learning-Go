package calculate

import (
	"CLI-Fin-Calc/services/arithmetic"
	"CLI-Fin-Calc/services/input"
	"fmt"
)

type PerformCalculation struct {
	Arithmetic arithmetic.ArithmeticService
	Input      input.UserInput
}

func (p PerformCalculation) Calculate() string {
	//init method instances
	arithmetic := arithmetic.ArithmeticService{}
	input := input.UserInput{}

	//user input
	fmt.Println("=========================\nPlease follow the following instructions")
	num, err := input.ReadNumbers()
	if err != nil {
		fmt.Println(err)
	}
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
	fmt.Println("\n========================================================================================")

	return continuance
}
