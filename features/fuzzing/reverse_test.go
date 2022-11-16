package fuzzing

import (
	"testing"
	"unicode/utf8"
)

// TestReverse tests the Reverse function.
func TestReverse(t *testing.T) {
	testcases := []struct {
		input    string
		expected string
	}{
		{"Hello, World!", "!dlroW ,olleH"},
		{" ", " "},
		{"!12345", "54321!"},
	}

	for _, tc := range testcases {
		actual := Reverse(tc.input)
		if actual != tc.expected {
			t.Errorf("Reverse(%s) = %s, expected %s", tc.input, actual, tc.expected)
		}
	}
}

// FuzzReverse is a fuzzing function for the Reverse function.
func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, World", " ", "!12345"}

	for _, tc := range testcases {
		f.Add(tc) // Add provides a new input to the fuzzer corpus.
	}

	// f.Fuzz runs the fuzzer for a while, and calls the given function
	// with each new input added with f.Add (orig == tc).
	f.Fuzz(func(t *testing.T, orig string) {
		rev := Reverse(orig)
		doubleRev := Reverse(rev)
		if doubleRev != orig {
			t.Errorf("Reverse(Reverse(%q)) = %q, expected %q", orig, doubleRev, orig)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string: %q", rev)
		}
	})
}
