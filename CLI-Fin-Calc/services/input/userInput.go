package input

import (
	"fmt"
)

type UserInput struct{}

func (n UserInput) ReadNumbers() ([2]int, error) {
	num1 := 0
	num2 := 0

	for {
		fmt.Println("\nEnter the first number")
		if _, err := fmt.Scan(&num1); err != nil {
			fmt.Println("Invalid Number Entered!! \nTry Again!")
			continue
		}
		break
	}

	for {
		fmt.Println("Enter the second number")
		if _, err := fmt.Scan(&num2); err != nil {
			fmt.Println("Invalid Number Entered!! \nTry Again!")
			continue
		}
		break
	}

	var numbersForOperation = [2]int{num1, num2}

	return numbersForOperation, nil
}

func (o UserInput) ReadOperand() string {
	oper := ""

	for {
		fmt.Println("Enter the operation to be performed: (+,-,*,/)")
		if _, err := fmt.Scan(&oper); err != nil {
			fmt.Println("Invalid Operand! \nTry Again!")
			continue
		}
		break
	}

	return oper
}
