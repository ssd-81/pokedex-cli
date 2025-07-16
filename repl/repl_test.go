package repl

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
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "is there  space in the sky",
			expected: []string{"is", "there", "space", "in", "the", "sky"},
		},
		{
			input:    "lost in the sky   ",
			expected: []string{"lost", "in", "the", "sky"},
		},
		{
			input:    "  test  of    less  ",
			expected: []string{"test", "of", "less"},
		},
	}

	for _, c := range cases {
		actual := CleanInput(strings.Join(strings.Fields(c.input), " ")) // returns a slice of strings

		if len(actual) != len(c.expected) {
			t.Errorf("length of actual and expected output is not equal; %d != %d", len(actual), len(c.expected))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("word mismatch between element at index %d", i)
			}
		}
	}
}
