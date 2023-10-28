package main

// //  hex complet 100%
// //  bin complet 100%
// //  up complet 100%
// //  low complet 100%
// //  cap complet 100%
// //  up or low or cap,num need to fix num > lengh 100%
// // func remove need to fix!!
// //  convert a to an 100%

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {

	// read the sample file
	fileContent, err := ioutil.ReadFile("sample.txt")
	if err != nil {
		log.Fatalf("Failed to read input file: %v", err)
	}

	// put it all to one variables or covart to string
	fileString := string(fileContent)

	// split the whole strings in to list
	list := strings.Split(fileString, " ")

	// call all functions
	err = findError(&list) // and find the errors
	findAletter(&list)     // convart a to an
	quote(&list)           // work on quote
	remove(&list)          // remove the space ...

	var str string
	if err != nil { // if there is error than str will be the error
		str = err.Error()
	} else { // other ways convart the list to one variable
		str = strings.Join(list, " ")
	}

	// chake the text file
	file, err := os.OpenFile("result.txt", os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Failed to open output file:", err)
		return
	}
	defer file.Close()

	// print the whole output to the file
	_, err = file.WriteString(str)
	if err != nil {
		fmt.Println("Failed to write to output file:", err)
		return
	}
}
