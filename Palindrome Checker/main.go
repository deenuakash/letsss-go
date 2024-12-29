package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Enter a string")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n') // to get input
	input = strings.TrimSpace(input)    // Trim \n from input

	//NOTE: Cannot use fmt.Scan/Scanf/Scanln as they consider space as a delimiter

	if Palindrome(input) {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Not Palindrome")
	}
}

func Palindrome(input string) bool {
	input = strings.ToLower(strings.ReplaceAll(input, " ", ""))
	left, right := 0, len(input)-1

	for left < right {
		if input[left] != input[right] {
			return false
		}
		left++
		right--
	}
	return true
}
