package menu

import "fmt"

type MainMenu struct{}

func (m MainMenu) Menu() int {
	option := 0

	// Calculator menu and start
	fmt.Println(`
	====================================
			Financial Calculator
	====================================

	1. Perform Calculation
	2. View Calculation History
	3. View Statistics
	4. Clear History
	5. Exit

	Select an option
	`)

	fmt.Scan(&option)

	return option
}
