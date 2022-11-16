package fuzzing

import (
	"fmt"
)

// Main is the entrypoint for the fuzzing package.
func Main() {
	input := "Hello, World!"
	rev := Reverse(input)
	doubleRev := Reverse(rev)
	fmt.Println("input:", input)
	fmt.Println("reverse:", rev)
	fmt.Println("double reverse:", doubleRev)
}

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
