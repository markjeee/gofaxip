package timeout


func CalculateAllottedTimeout(totalPages int) int {
	allottedTimeout := 400
	if totalPages >= 1 {
		allottedTimeout = allottedTimeout * totalPages
	}
	return allottedTimeout
}