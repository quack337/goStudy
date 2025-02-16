package main

import (
	"fmt"
)

func main3() { // slice
	var h = make([]int, 1, 2)		
	var a = h
	fmt.Println(len(a), cap(a), a)

	for i := 1; i < 70; i++ {
		h[0] = i
		a = append(a, i)
		fmt.Println(len(a), cap(a), a)
	}
}