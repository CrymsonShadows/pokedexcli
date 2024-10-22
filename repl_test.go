package main

import "testing"

func TestCleanUserInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    " ",
			expected: []string{},
		},
		{
			input:    "pokemon",
			expected: []string{"pokemon"},
		},
		{
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Hello World",
			expected: []string{"hello", "world"},
		},
		{
			input:    " Pokemon ",
			expected: []string{"pokemon"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Lengths are uneven between actual:'%v' vs expected:'%v'", actual, c.expected)
			continue
		}
		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("cleanInput(%s) == %v, expected: %v", c.input, actual[i], c.expected[i])
			}
		}
	}
}
