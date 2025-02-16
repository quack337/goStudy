package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	var s string = "ab가다ab"
	for _, ch := range s {
		fmt.Printf("%c, ", ch)
	}
	fmt.Printf("\n%d %d\n", len(s), utf8.RuneCountInString(s))

	var sa []rune = []rune("ab가다ab")
	for _, ch := range sa {
		fmt.Printf("%c, ", ch)
	}
	fmt.Printf("\n%d\n", len(sa))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%c, ", s[i])
	}
	fmt.Printf("\n%d\n", len(s))

	var sb []byte = []byte("ab가다ab")
	for _, ch := range sb {
		fmt.Printf("%c, ", ch)
	}
	fmt.Printf("\n%d\n", len(sb))

	var builder strings.Builder
	builder.WriteRune('A')
	builder.WriteString("BCD")
	fmt.Println(builder.String())
}