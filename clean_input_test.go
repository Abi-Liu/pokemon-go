package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{
			input:    " ",
			expected: "",
		},
		{
			input:    "Hello world",
			expected: "hello world",
		},
		{
			input:    "HELLO pokEmOn WorLD",
			expected: "hello pokemon world",
		},
		{
			input:    "hello",
			expected: "hello",
		},
	}

	for _, test := range cases {
		actual := cleanInput(test.input)
		if test.expected != actual {
			t.Errorf("Expected: %s\nReceived: %s", test.expected, actual)
		}
	}
}
