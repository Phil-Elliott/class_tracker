package utils

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{
			input:    "Hello world",
			expected: "hello",
		},
		{
			input:    "HELLO      ",
			expected: "hello",
		},
		{
			input:    "    ",
			expected: "",
		},
	}

	for _, cs := range cases {
		actual := CleanInput(cs.input)
		if actual != cs.expected {
			t.Errorf("The expected word %s is not equal to the result %s", cs.expected, actual)
		}

	}
}
