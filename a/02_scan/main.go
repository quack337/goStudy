package main

import (
	"fmt"
)

func main() {
	var a, b int
	var e error
	fmt.Print("a: ")
	_,e = fmt.Scanln(&a);					if e!=nil {goto E}
	fmt.Print("b: ")
	_,e = fmt.Scanln(&b);					if e!=nil {goto E}
	_,e = fmt.Printf("%T %T\n", a, b);		if e!=nil {goto E}
	_,e = fmt.Printf("%d %d\n", a, b);		if e!=nil {goto E}
	fmt.Print("a b: ")
	_,e = fmt.Scanln(&a, &b);				if e!=nil {goto E}
	_,e = fmt.Printf("%T %T\n", a, b);		if e!=nil {goto E}
	_,e = fmt.Printf("%d %d\n", a, b);		if e!=nil {goto E}
	return

E:
	fmt.Printf("%T, %s\n", e, e)
}
