package main

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func up(list *[]string) error {
	for i, word := range *list {

		// catch the up word
		pattern := regexp.MustCompile(`\(up`)
		if pattern.MatchString(word) { // if the word in list = up
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
					// convart the word to upper case
					for x := 1; x <= num; x++ {
						(*list)[i-x] = strings.ToUpper((*list)[i-x])
					}

					// remove the up word & the num
					removeAtIndex(list, i)
					removeAtIndex(list, i)
				}

			} else {

				// if it's (up) remove it else "up" leave it
				if pattern.MatchString(word) {
					if strings.HasSuffix(word, ")") {
						// remove the up word
						removeAtIndex(list, i)

						// convart to upper case
						(*list)[i-1] = strings.ToUpper((*list)[i-1])
					}
				}
			}
		}
	}
	return nil
}
