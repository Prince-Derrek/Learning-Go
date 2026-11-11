package main

import (
	calculate "CLI-Fin-Calc/services/calculator"
	"CLI-Fin-Calc/services/history"
	"CLI-Fin-Calc/services/input"
	"CLI-Fin-Calc/services/menu"
	"fmt"
)

func main() {
	//init services
	menuService := menu.MainMenu{}
	histService := history.NewHistoryService()
	calcService := calculate.PerformCalculation{
		HistoryService: histService,
	}
	userInput := input.UserInput{}

	//Action for each option
	option := 0
	for option != 5 {
		option = menuService.Menu()
		switch option {
		case 1:
			continuance := "y"

			for continuance == "y" {
				continuance = calcService.Calculate()
			}

		case 2:
			histOption := userInput.ReadSpecificHistory()
			switch histOption {
			case 1:
				additionHistory := histService.Get("addition")
				fmt.Printf("Addition History: \n %v", additionHistory)

			case 2:
				subtractionHistory := histService.Get("subtraction")
				fmt.Printf("Subtraction History: \n %v", subtractionHistory)

			case 3:
				productHistory := histService.Get("multiplication")
				fmt.Printf("Product History: \n %v", productHistory)

			case 4:
				divisionHistory := histService.Get("division")
				fmt.Printf("Division History: \n %v", divisionHistory)

			case 5:
				modulusHistory := histService.Get("modulus")
				fmt.Printf("Modulus History: \n %v", modulusHistory)

			case 6:
				percentageHistory := histService.Get("percentage")
				fmt.Printf("Percentage History: \n %v", percentageHistory)

			case 7:
				averageHistory := histService.Get("average")
				fmt.Printf("Average History: \n %v", averageHistory)

			case 8:
				minimumHistory := histService.Get("minimum")
				fmt.Printf("Minimum History: \n %v", minimumHistory)

			case 9:
				maximumHistory := histService.Get("maximum")
				fmt.Printf("Maximum History: \n %v", maximumHistory)

			case 10:
				allHistory := histService.GetAll()
				fmt.Printf("All Calculation History \n %v", allHistory)

			default:
				fmt.Println("Invalid Choice of Calculation Records")
			}
		}
	}
}
