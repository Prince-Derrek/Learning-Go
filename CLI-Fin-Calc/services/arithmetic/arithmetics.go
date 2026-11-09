package arithmetic

import (
	"math"
)

type ArithmeticService struct{}

func (s ArithmeticService) Sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total

}

func (s ArithmeticService) Diff(nums ...int) int {
	total := 0

	for _, num := range nums {
		total -= num
	}

	return total
}

func (s ArithmeticService) Product(nums ...int) int {
	total := 1

	for _, num := range nums {
		total *= num
	}

	return total
}

func (s ArithmeticService) Div(num1, num2 int) int {
	result := 0

	if num1 > num2 {
		result = num1 / num2
	} else {
		result = num2 / num1
	}

	return result
}

func (s ArithmeticService) Modulus(num1, num2 int) float64 {
	var no1 = float64(num1)
	var no2 = float64(num2)
	return math.Mod(no1, no2)
}

func (s ArithmeticService) Percentage(num1, num2 int) int {
	return (num1 / num2) * 100
}

func (s ArithmeticService) Average(nums []int) int {
	numberOfInputs := len(nums)
	if numberOfInputs == 0 {
		return 0
	}
	sum := 0
	for num := range nums {
		sum += (num)
	}

	return sum / numberOfInputs
}

func (s ArithmeticService) Minimum(nums []int) int {
	minimum := nums[0]

	for _, num := range nums {
		if num < minimum {
			minimum = num
		}
	}

	return minimum
}

func (s ArithmeticService) Maximum(nums []int) int {
	maximum := nums[0]

	for _, num := range nums {
		if num > maximum {
			maximum = num
		}
	}

	return maximum
}
