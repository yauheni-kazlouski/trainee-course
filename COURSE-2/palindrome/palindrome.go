package palindrome

import (
	"strings"
)

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