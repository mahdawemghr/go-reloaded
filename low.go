package main

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func low(list *[]string) error {
	for i, word := range *list {

		// catch the low word
		pattern := regexp.MustCompile(`\(low`)
		if pattern.MatchString(word) { // if the word in list = low
			if strings.HasSuffix(word, ",") { // if there is comma

				numPattern := regexp.MustCompile(`\d+`)
				match := numPattern.FindString((*list)[i+1])

				// convart string to intger
				num, err := strconv.Atoi(match)
				if err != nil {
					// handle error
				}

				// if the num bigger than the list
				if num > i {
					return errors.New("there is no " + strconv.Itoa(num) + " words")
				} else {
					// convart the word to lower case
					for x := 1; x <= num; x++ {
						(*list)[i-x] = strings.ToLower((*list)[i-x])
					}

					// remove the low word & the num
					removeAtIndex(list, i)
					removeAtIndex(list, i)
				}

			} else {

				// if it's (low) remove it else "low" leave it
				if pattern.MatchString(word) {
					if strings.HasSuffix(word, ")") {
						// remove the low word
						removeAtIndex(list, i)

						// convart to lower case
						(*list)[i-1] = strings.ToLower((*list)[i-1])
					}
				}
			}
		}
	}
	return nil
}
