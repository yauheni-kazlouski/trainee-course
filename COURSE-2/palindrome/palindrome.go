package palindrome

import (
	"strings"
)

func normalize(s string) string {
	temp := make([]rune, 0)

	for _, t := range(s) {
		if !strings.Contains(" !.,?", string(t)) {
			temp = append(temp, t)
		}
	}

	return string(temp)
}

// Compares runes of string one by one from both sides of given string
func HalfComparison(s string) bool {

	s = normalize(strings.ToLower(s))

	if len(s) == 0 {
		return false
	}

	rr := []rune(s)

	for i := 0; i < len(rr)/2; i++ {
		if rr[i] != rr[len(rr)-1-i] {
			return false
		}
	}

	return true
}


// Create a reverse string and compares it to the given one
func Reverse(s string) bool {

	s = normalize(strings.ToLower(s))

	if len(s) == 0 {
		return false
	}

	reversed := make([]rune, len(s))

	for i, t := range s {
		reversed[len(s) - 1 - i] = rune(t)
	}

	return string(reversed) == s
}


// Recursivly checks wether starting and ending symbols are equal until the edge cases
func doPalindrome(s []rune) bool {

	switch len(s){
	case 0:
		return false
	case 1:
		return true
	case 2:
		return s[0] == s[1]
	default:
		if s[0] != s[len(s)-1] {
			return false
		}
	
		return doPalindrome(s[1:len(s)-1])
	}
}

func Recursive(s string) bool {
	s = normalize(strings.ToLower(s))

	return doPalindrome([]rune(s))
}