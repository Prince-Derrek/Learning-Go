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

func (o UserInput) ReadOperand() int {
	oper := 0

	for {
		fmt.Println(`
		Enter the operation to be performed:
		1. Addition
		2. Subtraction
		3. Product
		4. Division
		5. Modulus
		6. Percentage Calculation
		7. Average of multiple numbers
		8. Minimum of multiple numbers
		9. Maximum of multiple numbers
		`)
		if _, err := fmt.Scan(&oper); err != nil {
			fmt.Println("Invalid Operand! \nTry Again!")
			continue
		}
		break
	}

	return oper
}

func (m UserInput) ReadMultipleIntegers() []int {
	numbersToOperate := 0
	fmt.Println("How many numbers would you like to enter:")
	fmt.Scan(&numbersToOperate)

	//Create a slice that can accept multiple inputs
	num := make([]int, 0, numbersToOperate)

	for i := 0; i < numbersToOperate; i++ {
		var number int

		fmt.Printf("Enter number %d: ", i+1)
		fmt.Scan(&number)

		num = append(num, number)
	}

	return num
}

func (h UserInput) ReadSpecificHistory() int {
	fmt.Println("Enter the History you would like to retrieve \n")

	history := 0

	for {
		fmt.Println(`
		Enter the history to be retrieved:
		1. Addition
		2. Subtraction
		3. Product
		4. Division
		5. Modulus
		6. Percentage Calculation
		7. Average of multiple numbers
		8. Minimum of multiple numbers
		9. Maximum of multiple numbers
		10. Retrieve All calculation history
		`)
		if _, err := fmt.Scan(&history); err != nil {
			fmt.Println("Invalid Operand! \nTry Again!")
			continue
		}
		break
	}

	return history
}

func (s UserInput) ReadStats() int {
	statOption := 0
	for {
		fmt.Println(`
		Which Statistic would you like:
		1. Total Calculations Performed
		2. Total Additions Performed
		3. Total Subtractions Performed
		4. Total Multiplications Performed
		5. Total Divisions Performed
		`)
		if _, err := fmt.Scan(&statOption); err != nil {
			fmt.Println("Invalid Option! \n Try Again!")
			continue
		}
		break
	}

	return statOption
}
