package calculate

import (
	"CLI-Fin-Calc/services/arithmetic"
	"CLI-Fin-Calc/services/history"
	"CLI-Fin-Calc/services/input"
	"fmt"
)

type PerformCalculation struct {
	Arithmetic     arithmetic.ArithmeticService
	Input          input.UserInput
	HistoryService *history.HistoryService
}

func (p PerformCalculation) Calculate() string {
	//init method instances
	arithmetic := arithmetic.ArithmeticService{}
	input := input.UserInput{}

	//user input
	fmt.Println("=========================\nPlease follow the following instructions")
	oper := input.ReadOperand()

	//operations
	switch oper {
	case 1:
		num, err := input.ReadNumbers()
		if err != nil {
			fmt.Println(err)
		}
		sum := arithmetic.Sum(num[0], num[1])
		fmt.Println("Sum:", sum)

		//Preserve Addition History
		p.HistoryService.Add("addition", history.CalculationRecord{
			Operands: []int{num[0], num[1]},
			Result:   float64(sum),
		})

	case 2:
		num, err := input.ReadNumbers()
		if err != nil {
			fmt.Println(err)
		}
		diff := arithmetic.Diff(num[0], num[1])
		fmt.Println("Difference: ", diff)

		//Preserve Subtraction Histort
		p.HistoryService.Add("subtraction", history.CalculationRecord{
			Operands: []int{num[0], num[1]},
			Result:   float64(diff),
		})

	case 3:
		num, err := input.ReadNumbers()
		if err != nil {
			fmt.Println(err)
		}
		prod := arithmetic.Product(num[0], num[1])
		fmt.Println("Product: ", prod)

		//Preserve Multiplication History
		p.HistoryService.Add("multiplication", history.CalculationRecord{
			Operands: []int{num[0], num[1]},
			Result:   float64(prod),
		})

	case 4:
		num, err := input.ReadNumbers()
		if err != nil {
			fmt.Println(err)
		}
		div := arithmetic.Div(float64(num[0]), float64(num[1]))
		fmt.Println("Quotient: ", div)

		//Preserve Division History
		p.HistoryService.Add("division", history.CalculationRecord{
			Operands: []int{num[0], num[1]},
			Result:   float64(div),
		})

	case 5:
		num, err := input.ReadNumbers()
		if err != nil {
			fmt.Println(err)
		}
		mod := arithmetic.Modulus(num[0], num[1])
		fmt.Printf("%v mod %v = %v", num[0], num[1], mod)

		//Preserve modulus history
		p.HistoryService.Add("modulus", history.CalculationRecord{
			Operands: []int{num[0], num[1]},
			Result:   float64(mod),
		})

	case 6:
		num, err := input.ReadNumbers()
		if err != nil {
			fmt.Println(err)
		}
		percent := arithmetic.Percentage(num[0], num[1])
		fmt.Printf("If %v is 100%% then %v is %v %% of %v ", num[0], num[1], percent, num[0])

		//Preserve Percentage History
		p.HistoryService.Add("percentage", history.CalculationRecord{
			Operands: []int{num[0], num[1]},
			Result:   float64(percent),
		})

	case 7:
		fmt.Println("Enter any number of integers")
		var numbers []int
		numbers = input.ReadMultipleIntegers()
		average := arithmetic.Average(numbers)
		fmt.Printf("The average of the numbers is: %v", average)

		//Preserve Avg History
		p.HistoryService.Add("average", history.CalculationRecord{
			Operands: numbers,
			Result:   float64(average),
		})

	case 8:
		fmt.Println("Enter any number of integers")
		numbers := input.ReadMultipleIntegers()
		minimum := arithmetic.Minimum(numbers)
		fmt.Printf("The minimum of the numbers is: %v", minimum)

		//Preserve Minimum history
		p.HistoryService.Add("minimum", history.CalculationRecord{
			Operands: numbers,
			Result:   float64(minimum),
		})

	case 9:
		fmt.Println("Enter any number of integers")
		numbers := input.ReadMultipleIntegers()
		maximum := arithmetic.Maximum(numbers)
		fmt.Printf("The maximum of the numbers is: %v", maximum)

		//Preserve maximum history
		p.HistoryService.Add("maximum", history.CalculationRecord{
			Operands: numbers,
			Result:   float64(maximum),
		})
	default:
		fmt.Println("Wrong operator input. Please use either of these (+,-,*,/)")
	}

	continuance := ""
	fmt.Println("\n\nWould you like to proceed (y if Yes, n if No)")
	fmt.Scan(&continuance)
	fmt.Println("\n========================================================================================5")

	return continuance
}
