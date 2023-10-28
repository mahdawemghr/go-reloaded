package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func hex(list *[]string) error {
	for i, word := range *list {
		pattern := regexp.MustCompile(`\(hex\)`) // catch the hex word
		if pattern.MatchString(word) {           // if the word in list = hex
			decimal, err := hexadecimalToDecimal((*list)[i-1])
			if err != nil {
				// Handle the error, e.g., log it or skip the replacement
				return err
			}
			(*list)[i-1] = decimal
			removeAtIndex(list, i) // remove the hex word
		}
	}
	return nil
}

func hexadecimalToDecimal(hex string) (string, error) {
	// Check for invalid characters in the input
	for i := range []byte(hex) {
		switch {
		case ('0' <= hex[i] && hex[i] <= '9') || ('A' <= hex[i] && hex[i] <= 'F') || ('a' <= hex[i] && hex[i] <= 'f'):
			continue
		default:
			return "", fmt.Errorf("Invalid input (hex)")
		}
	}

	decimalNum, err := strconv.ParseInt(hex, 16, 64)
	if err != nil {
		return "", err
	}

	return strconv.FormatInt(decimalNum, 10), nil
}
