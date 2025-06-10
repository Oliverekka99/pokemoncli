package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{"Hello World", []string{"hello", "world"}},
		{"  Leading and trailing spaces  ", []string{"leading", "and", "trailing", "spaces"}},
		{"", nil},
	}
	for _, cs := range cases {
		actual := cleanInput(cs.input)
		if len(actual) != len(cs.expected) {
			t.Errorf("The lenghts are not equal: %v , %v", len(actual), len(cs.expected))
			continue
		}
		for i := range actual {
			actualword := actual[i]
			expectedword := cs.expected[i]
			if actualword != expectedword {
				t.Errorf("Expected %s, but got %s", expectedword, actualword)
			}
		}
	}
}