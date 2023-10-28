package main

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func cap(list *[]string) error {
	for i, word := range *list {

		// catch the cap word
		pattern := regexp.MustCompile(`\(cap`)
		if pattern.MatchString(word) { // if the word in list = cap
			if strings.HasSuffix(word, ",") { // if there is comma
				numPattern := regexp.MustCompile(`\d+`)
				match := numPattern.FindString((*list)[i+1])

				num, err := strconv.Atoi(match)
				if err != nil {
					// handle error
				}

				// if the num bigger than the list
				if num > i {
					return errors.New("there is no " + strconv.Itoa(num) + " words")
				} else {
					// convart the first letter to upper case
					for x := 1; x <= num; x++ {
						(*list)[i-x] = strings.Title((*list)[i-x])
					}

					// remove te cap word & the num
					removeAtIndex(list, i)
					removeAtIndex(list, i)
				}

			} else {

				// if it's (cap) remove it else "cap" leave it
				if pattern.MatchString(word) {
					if strings.HasSuffix(word, ")") {
						// remove the cap word
						removeAtIndex(list, i)

						// convart the first letter to upper case
						(*list)[i-1] = strings.Title((*list)[i-1])
					}
				}
			}
		}
	}
	return nil
}
