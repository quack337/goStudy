package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i = 123
	var p1 *int = &i
	fmt.Println(*p1)

	var a = []int{1, 2, 3}
	var p2 *[]int = &a
	fmt.Println(*p2)
	fmt.Println(unsafe.Sizeof(a), unsafe.Sizeof(p2))

	type Point struct { x int; y int; }
	var p3 = &Point{1, 2}
	var p4 = new(Point)
	fmt.Println(*p3, *p4)
}