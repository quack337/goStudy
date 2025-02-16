package main

import "fmt"

func main() {
	var a, b = 2, 3
	fmt.Println(a, b)
	a, b = b, a
	fmt.Println(a, b)
}