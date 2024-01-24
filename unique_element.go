package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getUniqueElements(input string) ([]int, error) {
	// Remove "[" and "]" from the text array
	input = strings.Trim(input, "[]")

	// Split the text array by commas
	values := strings.Split(input, ",")

	// Convert the string values to integers and get unique elements
	uniqueMap := make(map[int]bool)
	uniqueElements := []int{}

	for _, value := range values {
		trimmedValue := strings.TrimSpace(value)
		num, err := strconv.Atoi(trimmedValue)
		if err != nil {
			return nil, err
		}

		if !uniqueMap[num] {
			uniqueMap[num] = true
			uniqueElements = append(uniqueElements, num)
		}
	}

	return uniqueElements, nil
}

func getUserInput() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter a array of integer: ")
	// Read input from the user
	scanner.Scan()
	// Get the user's input as a string
	userInput := scanner.Text()
	result, _ := getUniqueElements(userInput)
	fmt.Println("Array of Elements:", userInput)
	fmt.Println("Unique Elements: ", result)
}

func main() {
	getUserInput()
}
