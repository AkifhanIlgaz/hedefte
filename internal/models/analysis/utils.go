package analysis

func calculateAverage[T int | float64](oldAverage, newValue, itemCount T) T {
	return (oldAverage*(itemCount) + newValue) / (itemCount + 1)
}
