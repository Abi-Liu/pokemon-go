package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " ",
			expected: []string{},
		},
		{
			input:    "Hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "HELLO pokEmOn WorLD",
			expected: []string{"hello", "pokemon", "world"},
		},
		{
			input:    "hello",
			expected: []string{"hello"},
		},
	}

	for _, test := range cases {
		actual := cleanInput(test.input)
		if len(test.expected) != len(actual) {
			t.Errorf("Expected: %d\nReceived: %d", len(test.expected), len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := test.expected[i]
			if word != expectedWord {
				t.Errorf("cleanInput(%v) == %v, expected %v", test.input, actual, test.expected)
			}
		}
	}
}
