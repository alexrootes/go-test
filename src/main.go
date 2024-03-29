package main

import (
	"fmt"
//	"mypackage"
//	"strings"
//	"unicode/utf8"
	"time"
	"strconv"
)

var (
	j int64 = 0
)

func makeCakeAndSend(cs chan string) {
	j++
	cakeName := "Strawberry Cake " + strconv.FormatInt(j,10)
//	fmt.Println(strings.Title("Went to the shops"))
	fmt.Println("Making a cake and sending ...", cakeName)
	cs <- cakeName //send a strawberry cake
}

func receiveCakeAndPack(cs chan string) {
	for {
		s := <-cs //get whatever cake is on the channel
		fmt.Println("Packing received cake: ", s)
	}
}

func main() {
	cs := make(chan string)
	go receiveCakeAndPack(cs)
	for i := 0; i<100000; i++ {
		go makeCakeAndSend(cs)

		//sleep for a while so that the program doesn’t exit immediately and output is clear for illustration
		time.Sleep(10*time.Nanosecond)
	}
	time.Sleep(5*time.Second)
}

//func main() {
//	channelTest()
//	s := "Hello \u2129 world!"
//	s = mypackage.Clear(s)
//	fn := func(d rune) bool {
//		return strings.ContainsAny(string(d), " !")
//	}
//	for _, word := range strings.FieldsFunc(s, fn) {
//		fmt.Printf("|%v\n", word)
//		//fmt.Printf("|%v\n", idx)
//	}
//	bytes := []byte{88}
//	if utf8.FullRune(bytes) {
//		ru, _ := utf8.DecodeRune(bytes)
//		fmt.Printf("%c\n", ru)
//	} else {
//		fmt.Printf("not rune: %v\n", bytes)
//	}
//	for idx, _ := range s {
//		//fmt.Printf("%c", char)
//		r, _ := utf8.DecodeRuneInString(s[idx:])
//		fmt.Printf("%c", r)
//	}
//}

//func fieldFunc(d rune) bool {
//	return strings.ContainsAny(string(d), " !")
//}
