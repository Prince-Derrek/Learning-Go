package main

import (
	calculate "CLI-Fin-Calc/services/calculator"
	"CLI-Fin-Calc/services/history"
	"CLI-Fin-Calc/services/input"
	"CLI-Fin-Calc/services/menu"
	"CLI-Fin-Calc/services/statistics"
	"fmt"
)

func main() {
	//init services
	menuService := menu.MainMenu{}
	histService := history.NewHistoryService()
	calcService := calculate.PerformCalculation{
		HistoryService: histService,
	}
	statService := statistics.CalculationStatistics{
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
				additionHistory := histService.Get(history.Addition)
				fmt.Printf("Addition History: \n %v", additionHistory)

			case 2:
				subtractionHistory := histService.Get(history.Subtraction)
				fmt.Printf("Subtraction History: \n %v", subtractionHistory)

			case 3:
				productHistory := histService.Get(history.Multiplication)
				fmt.Printf("Product History: \n %v", productHistory)

			case 4:
				divisionHistory := histService.Get(history.Division)
				fmt.Printf("Division History: \n %v", divisionHistory)

			case 5:
				modulusHistory := histService.Get(history.Modulus)
				fmt.Printf("Modulus History: \n %v", modulusHistory)

			case 6:
				percentageHistory := histService.Get(history.Percentage)
				fmt.Printf("Percentage History: \n %v", percentageHistory)

			case 7:
				averageHistory := histService.Get(history.Average)
				fmt.Printf("Average History: \n %v", averageHistory)

			case 8:
				minimumHistory := histService.Get(history.Minimum)
				fmt.Printf("Minimum History: \n %v", minimumHistory)

			case 9:
				maximumHistory := histService.Get(history.Maximum)
				fmt.Printf("Maximum History: \n %v", maximumHistory)

			case 10:
				allHistory := histService.GetAll()
				fmt.Printf("All Calculation History \n %v", allHistory)

			default:
				fmt.Println("Invalid Choice of Calculation Records")
			}

		case 3:
			statOption := userInput.ReadStats()
			switch statOption {
			case 1:
				totalCalculations := statService.Total()
				fmt.Printf("Total Calculations Performed = %v", totalCalculations)
			case 2:
				totalAdditions := statService.SpecificTotals(history.Addition)
				fmt.Printf("Total Additions Performed = %v", totalAdditions)
			case 3:
				totalSubtractions := statService.SpecificTotals(history.Subtraction)
				fmt.Printf("Total Subtractions Performed = %v", totalSubtractions)
			case 4:
				totalMultiplications := statService.SpecificTotals(history.Multiplication)
				fmt.Printf("Total Multiplications Performed = %v", totalMultiplications)
			case 5:
				totalDivisions := statService.SpecificTotals(history.Division)
				fmt.Printf("Total Divisions Performed = %v", totalDivisions)
			case 6:
				largestResult := statService.LargestResult()
				fmt.Printf("The Largest result obtained is : %v", largestResult)
			default:
				fmt.Println("Invalid Option for Statistics")
			}

		}
	}
}
