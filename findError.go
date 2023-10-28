package main

func findError(list *[]string) error {
	if err := cap(list); err != nil { // first letter in word to upper case
		return err
	}

	if err := up(list); err != nil { // whole word to upper case
		return err
	}

	if err := low(list); err != nil { // whole word to lower case
		return err
	}

	if err := bin(list); err != nil { // convart the hexadecimal number
		return err
	}

	if err := hex(list); err != nil { // convart the binary number
		return err
	}

	// ... continue with the rest of the method logic ...

	return nil
}
