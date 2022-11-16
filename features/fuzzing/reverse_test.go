package fuzzing

import (
	"testing"
)

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
