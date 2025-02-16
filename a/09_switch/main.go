package main

import "fmt"

func main() {
	var a = 2

	var s string
	switch a {
	case 0: s = "zero"
	case 1: s = "one"
	case 2: s = "two"
	}
	fmt.Println(s)

	switch b := 2; b {
	case 0: s = "zero"
	case 1: s = "one"
	case 2: s = "two"
	}
	fmt.Println(s)

	switch true {
	case a == 0: s = "zero"
	case a == 1: s = "one"
	case a == 2: s = "two"
	}
	fmt.Println(s)

	a = 0
	switch a {
	case 0: 
		fmt.Println("zero")
		fallthrough
	case 1: 
		fmt.Println("one")
		fallthrough
	case 2: 
		fmt.Println("two")
	}
}