package main

import (
	calculate "CLI-Fin-Calc/services/calculator"
	"CLI-Fin-Calc/services/menu"
)

func main() {
	//init services
	menuService := menu.MainMenu{}
	calcService := calculate.PerformCalculation{}
	option := 0
	for option != 5 {
		option = menuService.Menu()
		switch option {
		case 1:
			continuance := "y"

			for continuance == "y" {
				continuance = calcService.Calculate()
			}
		}
	}
}
