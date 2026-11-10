package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{{
		input:    "  hello  world  ",
		expected: []string{"hello", "world"},
	}}

	for _, c := range cases {
		got := cleanInput(c.input)

		if len(got) != len(c.expected) {
			t.Errorf("cleanInput(%q) == %q, want %q", c.input, got, c.expected)
			continue // avoid indexing past the end below
		}

		for i, e := range got {
			if e != c.expected[i] {
				t.Errorf("cleanInput(%q) == %q, want %q", c.input, got, c.expected)
			}
		}
	}

}
