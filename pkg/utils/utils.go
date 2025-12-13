package utils

import "time"

func CalculateTotalPages(totalCount, limit int) int {
	if limit <= 0 {
		return 0
	}

	return (totalCount + limit - 1) / limit
}

func GetStart(timeInterval int) time.Time {
	switch timeInterval {
	case 1:
		return time.Now().AddDate(0, -1, 0)
	case 3:
		return time.Now().AddDate(0, -3, 0)
	case 6:
		return time.Now().AddDate(0, -6, 0)
	default:
		return time.Now().AddDate(-2, 0, 0)
	}
}
