package fuzzing

// Reverse returns the reverse of the given string.
func Reverse(s string) string {
	r := []rune(s) // convert string to rune slice

	// i start from 0, j start from the last index of the slice
	// i and j are incremented and decremented respectively
	// until i is equal to the half of the length of the slice
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i] // swap
	}
	return string(r)
}
