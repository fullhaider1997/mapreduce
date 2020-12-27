package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {

	content := "Haider .  Ibrahim  3 ahmed  alhafiz"

	ff := func(r rune) bool { return !unicode.IsLetter(r) }

	words := strings.FieldsFunc(content, ff)

	for indx, val := range words {

		fmt.Println(indx, val)
	}

	fmt.Println(len(words))

}
