package services

import "fmt"

type UserInput struct{}

func (n UserInput) Variant() [2]int {
	num1 := 0
	num2 := 0
	fmt.Println("Enter the first number")
	fmt.Scan(&num1)

	fmt.Println("Enter the second number")
	fmt.Scan(&num2)

	var numbersForOperation = [2]int{num1, num2}

	return numbersForOperation
}

func (o UserInput) Operand() string {
	oper := ""

	fmt.Println("Enter the operation to be performed: (+,-,*,/)")
	fmt.Scan(&oper)

	return oper
}
