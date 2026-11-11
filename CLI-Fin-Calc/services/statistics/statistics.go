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
func (s CalculationStatistics) LargestResult() (float64, bool) {
	totalHistory := s.HistoryService.GetAll()

	var largest float64
	found := false

	for _, records := range totalHistory {
		for _, record := range records {
			if !found || record.Result > largest {
				largest = record.Result
				found = true
			}
		}
	}

	return largest, found
}

func (s CalculationStatistics) SmallestResult() (float64, bool) {
	totalHistory := s.HistoryService.GetAll()

	var smallest float64
	found := false

	for _, records := range totalHistory {
		for _, record := range records {
			if !found || record.Result < smallest {
				smallest = record.Result
				found = true
			}
		}
	}

	return smallest, found
}

func (s CalculationStatistics) AverageResult() float64 {
	totalHistory := s.HistoryService.GetAll()

	summation := 0.0

	for _, records := range totalHistory {
		for _, record := range records {
			summation += record.Result
		}
	}

	noOfOps := len(totalHistory)

	average := summation / float64(noOfOps)

	return average
}
