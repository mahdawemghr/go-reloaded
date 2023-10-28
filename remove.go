package main

import (
	"regexp"
	"strings"
)

func remove(list *[]string) {
	for i := 0; i < len(*list); i++ {
		word := (*list)[i]
		if punctuations(string(word)) {
			containsLetter := regexp.MustCompile(`[a-z]`).MatchString(word)
			if containsLetter {
				first := word[0]

				(*list)[i-1] += string(first)
				(*list)[i] = strings.ReplaceAll((*list)[i], string(first), "")
			} else {
				(*list)[i-1] = (*list)[i-1] + word

				removeAtIndex(list, i)
			}
		}
	}
}

func punctuations(word string) bool {
	if len(word) > 0 {
		firstLetter := word[0]
		switch firstLetter {
		case '.', ',', '!', '?', ':', ';':
			return true
		}
	}
	return false
}

func quote(list *[]string) {
	count := 0
	for i := 0; i < len(*list); i++ {
		if (*list)[i] == "'" {
			if count%2 == 0 {
				count++
				if i+1 < len(*list) {
					(*list)[i+1] = (*list)[i] + (*list)[i+1]
				}
				removeAtIndex(list, i)
			} else {
				if i > 0 {
					(*list)[i-1] = (*list)[i-1] + (*list)[i]
				}
				removeAtIndex(list, i)
				i-- // Decrement i to account for the removed element
				count--
			}
		}
	}
}
