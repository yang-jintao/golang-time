package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("ss")
	s := "小比分"
	//m := []rune(s)
	fmt.Println(utf8.RuneCountInString(s))
}
