package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "i used to be an adventurer like you",
			expected: []string{"i", "used", "to", "be", "an", "adventurer", "like", "you"},
		},
		{
			input: "",
			expected: []string{},
		},
		{
			input: "owo",
			expected: []string{"owo"},
		},
		{
			input: "But then I took an arrow to the knee",
			expected: []string{"but", "then", "i", "took", "an", "arrow", "to", "the", "knee"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected){
			t.Errorf("test failed: got %s expected %s", actual, c.expected)
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("test failed: got %s expected %s", word, c.expected)
				continue
			}
		}
	}
}