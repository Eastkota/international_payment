package helpers

// DollarsToCents converts a dollar amount to cents
func DollarsToCents(dollars float64) int64 {
	return int64(dollars * 100)
}