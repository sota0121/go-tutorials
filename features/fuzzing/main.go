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
