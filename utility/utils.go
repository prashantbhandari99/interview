package utility

func Add(a, b float64) float64 {
	return a + b
}

func Subtract(a, b float64) float64 {
	if a > b {
		return a - b
	}
	return b - a
}
