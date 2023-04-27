package math

func Sum(a [5]int) int {
	result := 0
	for _, value := range a {
		result += value
	}
	return result
}
