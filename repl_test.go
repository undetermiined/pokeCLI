package main

import (
	"reflect"
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

		if !reflect.DeepEqual(actual, c.expected) {
			t.Errorf("cleanInput(%q): expected %v, got %v", c.input, c.expected, actual)
		}
	}
}
