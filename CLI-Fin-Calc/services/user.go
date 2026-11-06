package services

import (
	"errors"
	"fmt"
)

type UserInput struct{}

func (n UserInput) ReadNumbers() ([2]int, error) {
	num1 := 0
	num2 := 0

	for {
		fmt.Println("Enter the first number")
		if _, err := fmt.Scan(&num1); err != nil {
			return [2]int{}, errors.New("Invalid first number!")
		}
		break
	}

	for {
		fmt.Println("Enter the second number")
		if _, err := fmt.Scan(&num2); err != nil {
			return [2]int{}, errors.New("Invalid first number!")
		}
		break
	}

	var numbersForOperation = [2]int{num1, num2}

	return numbersForOperation, nil
}

func (o UserInput) ReadOperand() string {
	oper := ""

	fmt.Println("Enter the operation to be performed: (+,-,*,/)")
	fmt.Scan(&oper)

	return oper
}
