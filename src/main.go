package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

//func fieldFunc(d rune) bool {
//	return strings.ContainsAny(string(d), " !")
//}

func main() {
	s := "Hello \u2129 world!"
	fmt.Println(s)
	fn := func(d rune) bool {
		return strings.ContainsAny(string(d), " !")
	}
	for _, word := range(strings.FieldsFunc(s, fn)) {
		fmt.Printf("|%v\n", word)
		//fmt.Printf("|%v\n", idx)
	}
	for idx, _ := range(s) {
		//fmt.Printf("%c", char)
		r, _ := utf8.DecodeRuneInString(s[idx:])
		fmt.Printf("%c", r)
	}
}
