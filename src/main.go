package main

import (
	"fmt"
	"mypackage"
	"strings"
	"unicode/utf8"
)

func main() {
	s := "Hello \u2129 world!"
	s = mypackage.Clear(s)
	fn := func(d rune) bool {
		return strings.ContainsAny(string(d), " !")
	}
	for _, word := range strings.FieldsFunc(s, fn) {
		fmt.Printf("|%v\n", word)
		//fmt.Printf("|%v\n", idx)
	}
	bytes := []byte{88}
	if utf8.FullRune(bytes) {
		ru, _ := utf8.DecodeRune(bytes)
		fmt.Printf("%c\n", ru)
	} else {
		fmt.Printf("not rune: %v\n", bytes)
	}
	for idx, _ := range s {
		//fmt.Printf("%c", char)
		r, _ := utf8.DecodeRuneInString(s[idx:])
		fmt.Printf("%c", r)
	}
}

//func fieldFunc(d rune) bool {
//	return strings.ContainsAny(string(d), " !")
//}
