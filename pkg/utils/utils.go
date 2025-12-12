package utils

func CalculateTotalPages(totalCount, limit int) int {
	if limit <= 0 {
		return 0
	}

	return (totalCount + limit - 1) / limit
}
