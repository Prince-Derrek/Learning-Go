package statistics

import (
	"CLI-Fin-Calc/services/history"
)

type CalculationStatistics struct {
	HistoryService *history.HistoryService
}

/*
Calculate:
Total Calculations Performed
Number of additions, subtractions, multiplication and divisions
Largest result
Smallest result
Average of all recorded results
*/

//Calculating total calculations performed
func (s CalculationStatistics) Total() int {
	totalHistory := s.HistoryService.GetAll()
	totalCalculations := len(totalHistory)

	return totalCalculations
}


//Calculating Specific operation calculations
func (s CalculationStatistics) SpecificTotals(operation string) int {
	specificHistory := s.HistoryService.Get(operation)
	specificCalculations := len(specificHistory)

	return specificCalculations
}

//Calculating largest result


