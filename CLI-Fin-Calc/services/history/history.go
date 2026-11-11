package history

type CalculationRecord struct {
	Operands []int
	Result   float64
}

type HistoryService struct {
	records map[string][]CalculationRecord
}

// Constructor
func NewHistoryService() *HistoryService {
	return &HistoryService{
		records: make(map[string][]CalculationRecord),
	}
}

// Store a Calculataion
func (h *HistoryService) Add(
	operation string,
	record CalculationRecord,
) {
	h.records[operation] = append(
		h.records[operation],
		record,
	)
}

// Retrieve calculations for a particular operation
func (h *HistoryService) Get(operation string) []CalculationRecord {
	return h.records[operation]
}

// Retrieve all calculations
func (h *HistoryService) GetAll() map[string][]CalculationRecord {
	return h.records
}

// Clear all history
func (h *HistoryService) Clear() {
	clear(h.records)
}
