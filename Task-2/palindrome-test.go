package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Simple palindrome",
			input:    "madam",
			expected: true,
		},
		{
			name:     "Mixed case palindrome",
			input:    "RaceCar",
			expected: true,
		},
		{
			name:     "Palindrome with punctuation",
			input:    "A man, a plan, a canal: Panama!",
			expected: true,
		},
		{
			name:     "Not a palindrome",
			input:    "Hello world",
			expected: false,
		},
		{
			name:     "Empty string",
			input:    "",
			expected: true,
		},
		{
			name:     "Single character",
			input:    "a",
			expected: true,
		},
		{
			name:     "Palindrome with numbers",
			input:    "12321",
			expected: true,
		},
		{
			name:     "Not a palindrome with numbers",
			input:    "12345",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsPalindrome(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %v, but got %v", tt.expected, result)
			}
		})
	}
}
