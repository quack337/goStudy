package main

import "fmt"

const (
	A = iota
	B 
	C
)

func main() {
	fmt.Println(A, B, C)
	fmt.Printf("%T %d\n", A, A)
}
