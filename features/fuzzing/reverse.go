package fuzzing

// Reverse returns the reverse of the given string.
func Reverse(s string) string {
	b := []byte(s)

	// i start from 0, j start from the last index of the slice
	// i and j are incremented and decremented respectively
	// until i is equal to the half of the length of the slice
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i] // swap
	}
	return string(b)
}
