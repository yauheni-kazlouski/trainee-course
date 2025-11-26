package reverse

import (
	//"regexp"
	"strings"
)

func ClearReverse(s string) string {
	words := make([][]rune, 0)
	// Splits a string into a slice of words
	temp := make([]rune, 0)
	for _, t := range s {
		if t == ' ' {
			if len(temp) == 0 {
				continue
			}
			words = append(words, temp)
			temp = make([]rune, 0)
			continue
		}

		temp = append(temp, t)
	}

	if len(temp) != 0 {
		words = append(words, temp)
	}
    
	// If there is no words, return given string
	if len(words) == 0 {
		return s
	}

	// reverse a word by going from both sides of it and exchanging i and len - i elements
	for _, word := range words {
		for i := 0; i < len(word)/2; i++ {
			word[i], word[len(word)-1-i] = word[len(word)-1-i], word[i]
		}
	}

	reversed := string(words[0])

	for i, word := range words {
		if i == 0 {
			continue
		}
		reversed += " " + string(word)
	}

	return reversed
}


func reverseWord(s string) string{
	temp := []byte(s)

	for i := 0; i < len(temp)/2; i++ {
		temp[i], temp[len(temp)-1-i] = temp[len(temp)-1-i], temp[i]
	}

	return string(temp)
}

func PoweredReverse(s string) string {
	words := strings.Fields(s);

	if len(words) == 0{
		return s
	}

	markupStr := strings.Replace(s, words[0], "%s", 1)
	words[0] = reverseWord(words[0])
	for i := 1; i < len(words); i++{
		markupStr = strings.Replace(markupStr, words[i], "%s", 1)
		words[i] = reverseWord(words[i])
	}

	return strings.Join(words, " ")	
}