package main

import (
	"strings"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  hello,  world  ",
			expected: []string{"hello,", "world"},
		},
		{
			input:    "  --hello - world  ",
			expected: []string{"--hello", "-", "world"},
		},
		{
			input:    "  hello  world !  ",
			expected: []string{"hello", "world", "!"},
		},
	}
	for _, c := range cases {
		actualInput := cleanInput(c.input)
		if len(actualInput) != len(c.expected) {
			t.Errorf("length of %v:%d does not match %v:%d", actualInput, len(actualInput), c.expected, len(c.expected))
			t.Fail()
		}
		for i := range actualInput {
			word := (actualInput[i])
			expectedWord := (c.expected[i])
			if !strings.EqualFold(word, expectedWord) {
				t.Errorf("Word %s does not match %s", word, expectedWord)
				t.Fail()
			}
		}
	}
}
