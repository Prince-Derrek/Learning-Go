package services

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
