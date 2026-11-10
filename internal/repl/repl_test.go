package repl

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "HELP",
			expected: []string{"help"},
		},
		{
			input:    "Catch Pikachu",
			expected: []string{"catch", "pikachu"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "   \t\n ",
			expected: []string{},
		},
	}

	for _, test := range tests {
		got := cleanInput(test.input)
		if !reflect.DeepEqual(got, test.expected) {
			t.Errorf("cleanInput(%q) == %q, want %q", test.input, got, test.expected)
		}
	}
}

func TestGetCommandsContainsAllCommands(t *testing.T) {
	want := []string{"help", "map", "mapb", "explore", "catch", "inspect", "pokedex", "exit"}

	commands := getCommands()
	for _, name := range want {
		if _, ok := commands[name]; !ok {
			t.Errorf("expected command %q to be registered", name)
		}
	}
	if len(commands) != len(want) {
		t.Errorf("got %d commands, want %d", len(commands), len(want))
	}
}

func TestSortedCommandsIsOrdered(t *testing.T) {
	commands := sortedCommands()
	for i := 1; i < len(commands); i++ {
		if commands[i-1].name > commands[i].name {
			t.Fatalf("commands are not sorted: %q before %q", commands[i-1].name, commands[i].name)
		}
	}
}
