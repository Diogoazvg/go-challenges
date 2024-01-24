package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getUserInput() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your text: ")
	// Read input from the user
	scanner.Scan()
	// Get the user's input as a string
	userInput := scanner.Text()
	// fmt.Printf("%s is a palindrome: \n", isPalindrome(userInput))
	fmt.Printf("%s is a palindrome: %t\n", userInput, isPalindrome(userInput))
}

// isPalindrome checks if a given string is a palindrome.
func isPalindrome(s string) bool {
	// Convert the string to lowercase and remove spaces
	s = strings.ToLower(strings.ReplaceAll(s, " ", ""))

	// Compare characters from both ends of the string
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func main() {
	getUserInput()
}

// Example palindromes
// palindrome1 := "A man a plan a canal Panama"
// palindrome2 := "radar"

// // Example non-palindrome
// nonPalindrome := "hello"
