package main

import (
	"reflect"
	"testing"
)


func TestWordFrequencyCount(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string]int
	}{
		{
			name:  "Simple case",
			input: "Hello world",
			expected: map[string]int{
				"hello": 1,
				"world": 1,
			},
		},
		{
			name:  "Case insensitivity",
			input: "Hello hello HeLLo",
			expected: map[string]int{
				"hello": 3,
			},
		},
		{
			name:  "Punctuation handling",
			input: "Hello, world! It's a wonderful world.",
			expected: map[string]int{
				"hello":     1,
				"world":     2,
				"it's":      1,
				"a":         1,
				"wonderful": 1,
			},
		},
		{
			name:  "Mixed case with numbers",
			input: "123 go 123 Go!",
			expected: map[string]int{
				"123": 2,
				"go":  2,
			},
		},
		{
			name:     "Empty input",
			input:    "",
			expected: map[string]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := WordFrequencyCount(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v, but got %v", tt.expected, result)
			}
		})
	}
}
