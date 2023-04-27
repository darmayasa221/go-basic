package math

func Iteration(c string) string {
	var result string
	for i := 0; i < 5; i++ {
		result += c
	}
	return result
}
