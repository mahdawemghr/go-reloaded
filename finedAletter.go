package main

func findAletter(list *[]string) {
	for i := 0; i < len(*list)-1; i++ {
		if (*list)[i] == "a" && isVowel((*list)[i+1]) {
			(*list)[i] = "an"
		}
	}
}

func isVowel(word string) bool {
	if len(word) > 0 {
		firstLetter := word[0]
		switch firstLetter {
		case 'a', 'e', 'i', 'o', 'u', 'h':
			return true
		}
	}
	return false
}
