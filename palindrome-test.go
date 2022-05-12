package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	var strInput string
	if len(os.Args) > 2 {
		strInput = strings.Join(os.Args[1:], " ")
	} else {
		strInput = os.Args[1]
	}

	// clean the string for uppercase, space and another sepcial character
	// to support multiple words
	strInput = cleanString(strings.ToLower(strInput))

	fmt.Println(checkPalindrome(strInput))
}

// function to check palindrome
func checkPalindrome(words string) bool {
	count := len(words)

	// loop for half the word, because we compare half of the string
	for i := 0; i < count/2; i++ {
		if words[i] != words[count-1-i] {
			return false
		}
	}
	return true
}

// function to clean sign : . , space -
func cleanString(sentence string) string {
	re := regexp.MustCompile("[" + "., -" + "]+")
	result := re.ReplaceAllString(sentence, "")
	return result
}
