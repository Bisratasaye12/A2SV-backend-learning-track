package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)


func IsPalindrome(input string) bool {
	input = strings.ToLower(input)

	isLetterOrDigit := func(r rune) bool {
		return unicode.IsLetter(r) || unicode.IsDigit(r)
	}

	var filteredInput []rune
	for _, r := range input {
		if isLetterOrDigit(r) {
			filteredInput = append(filteredInput, r)
		}
	}

	for i, j := 0, len(filteredInput)-1; i < j; i, j = i+1, j-1 {
		if filteredInput[i] != filteredInput[j] {
			return false
		}
	}

	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a string: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if IsPalindrome(input) {
		fmt.Println("The string is a palindrome.")
	} else {
		fmt.Println("The string is not a palindrome.")
	}
}
