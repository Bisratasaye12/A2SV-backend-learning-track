package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func WordFrequencyCount(input string) map[string]int {
	input = strings.ToLower(input)

	isLetterOrDigit := func(r rune) bool {
		return unicode.IsLetter(r) || unicode.IsDigit(r)
	}
	
	var words []string
	var curr = strings.Fields(input)
	for _, w := range curr{
        cleaned := strings.FieldsFunc(w, func(r rune) bool {
		return !isLetterOrDigit(r) && !unicode.IsSpace(r)
     	})
     	
	    words = append(words,cleaned...)
	    fmt.Println(cleaned)
	    
	}
    
	frequency := make(map[string]int)

	for _, word := range words {
		frequency[word]++
	}

	return frequency
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a string: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	frequency := WordFrequencyCount(input)

	fmt.Println("Word Frequency Count:")
	for word, count := range frequency {
		fmt.Printf("%s: %d\n", word, count)
	}
}
