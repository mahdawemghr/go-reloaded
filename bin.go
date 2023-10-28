package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func bin(list *[]string) error {
	for i, word := range *list {
		pattern := regexp.MustCompile(`\(bin\)`) // catch the bin word
		if pattern.MatchString(word) {           // if the word in list = bin
			numPattern := regexp.MustCompile(`\d+`)
			match := numPattern.FindString((*list)[i-1])
			decimal, err := binaryToDecimal(match)
			if err != nil {
				// Handle the error, e.g., log it or skip the replacement
				return err
			}
			(*list)[i-1] = decimal
			removeAtIndex(list, i) // remove the bin word
		}
	}
	return nil
}

func binaryToDecimal(bin string) (string, error) {
	decimal, err := strconv.ParseInt(bin, 2, 8)
	if err != nil {
		return "", fmt.Errorf("Invalid input (bin)")
	}
	return strconv.FormatInt(decimal, 10), nil
}
