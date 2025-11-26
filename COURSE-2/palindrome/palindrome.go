package palindrome

import (
	"strings"
)

// Compares runes of string one by one from both sides of given string
func HalfComparison(s string) bool {

	s = strings.ToLower(s)
	
	temp := make([]rune, 0)

	for _, t := range(s) {
		if !strings.Contains(" !.,?", string(t)) {
			temp = append(temp, t)
		}
	}

	s = string(temp)

	if len(s) == 0 {
		return false
	}

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}

	return true
}


// Create a reverse string and compares it to the given one
func Reverse(s string) bool {

	s = strings.ToLower(s)
	
	temp := make([]rune, 0)

	for _, t := range(s) {
		if !strings.Contains(" !.,?", string(t)) {
			temp = append(temp, t)
		}
	}

	s = string(temp)

	if len(s) == 0 {
		return false
	}

	reversed := make([]byte, len(s))

	for i, t := range s {
		reversed[len(s)- 1 - i] = byte(t)
	}

	return string(reversed) == s
}


// Recursivly checks wether starting and ending symbols are equal until the edge cases
func doPalindrome(s string) bool {

	if len(s) == 0{
		return false
	}

	if len(s) == 1 {
		return true
	}
	if len(s) == 2{
		return s[0] == s[1]
	}

	if s[0] != s[len(s)-1] {
		return false
	}
	
	return doPalindrome(s[1:len(s)-1])
}

func Recursive(s string) bool {
	s = strings.ToLower(s)
	
	temp := make([]rune, 0)

	for _, t := range(s) {
		if !strings.Contains(" !.,?", string(t)) {
			temp = append(temp, t)
		}
	}

	s = string(temp)

	return doPalindrome(s)
}