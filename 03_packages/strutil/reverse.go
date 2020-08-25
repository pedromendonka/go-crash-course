package strutil

// Reverse a string
func Reverse(s string) string {
	r := []rune(s)
	var result []rune
	for i := len(r) - 1; i >= 0; i-- {
		result = append(result, r[i])
	}
	return string(result)
}
