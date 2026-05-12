package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " Snorlax eEvEe   bulbaSauR",
			expected: []string{"snorlax", "eevee", "bulbasaur"},
		},
		{
			input:    "GYARADOs  liTTen Espeon  ",
			expected: []string{"gyarados", "litten", "espeon"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		//	Result length comparison
		if len(actual) != len(c.expected) {
			t.Errorf(`---------------------------------
Input:     (%v)
Expecting:  %v
Actual:     %v
Fail
`, c.input, c.expected, actual)
		}

		//	Result word comparison
		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf(`---------------------------------
Input:     (%v)
Expecting:  %v
Actual:     %v
Fail
`, c.input, c.expected, actual)
			}
		}

		fmt.Printf(`---------------------------------
Input:     (%v)
Expecting:  %v
Actual:     %v
Pass!
`, c.input, c.expected, actual)
	}
}
