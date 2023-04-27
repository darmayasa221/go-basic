package math

func ArraySlice(a []int) int {
	var result int
	for _, value := range a {
		result += value
	}
	return result
}
