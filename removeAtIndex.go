package main

func removeAtIndex(list *[]string, index int) {
	// func to remove the index I want
	*list = append((*list)[:index], (*list)[index+1:]...)
}
