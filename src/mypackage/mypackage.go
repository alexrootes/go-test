package mypackage

import "strings"

func Clear(word string) string {
	if strings.HasPrefix(word, "Hello") {
		return strings.Replace(word, "Hello", "World", -1)
	}
	return word
}
