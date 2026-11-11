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

// Calculating total calculations performed
func (s CalculationStatistics) Total() int {
	totalHistory := s.HistoryService.GetAll()
	totalCalculations := len(totalHistory)

	return totalCalculations
}

// Calculating Specific operation calculations
func (s CalculationStatistics) SpecificTotals(operation string) int {
	specificHistory := s.HistoryService.Get(operation)
	specificCalculations := len(specificHistory)

	return specificCalculations
}

// Calculating largest result
func (s CalculationStatistics) LargestResult() float64 {
	totalHistory := s.HistoryService.GetAll()

	largest := 0.0

	for _, records := range totalHistory {
		for _, record := range records {
			if record.Result > largest {
				largest = record.Result
			}
		}
	}

	return largest
}
