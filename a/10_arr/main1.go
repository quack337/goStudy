package main

import (
	"fmt"
	"unsafe"
)

func main1() { // array
	var a = [4]int {1, 2, 3, 4}
	var b = [3]int {11, 22, 33}
	var c = [2]int {111, 222}

	fmt.Println(unsafe.Sizeof(a), unsafe.Sizeof(b), unsafe.Sizeof(c))
	fmt.Println(a, b, c)

	var d = [2][3]int {{1, 2, 3}, {4, 5, 6}}

	fmt.Println(unsafe.Sizeof(d))
	fmt.Println(d)

	d[1] = d[0]
	d[0][0] = -11
	fmt.Println(d)

	fmt.Println(unsafe.Sizeof(d))

	var e = [3]int{1, 2, 3}
	f2(e) // call by value
	fmt.Println(e)
}

func f2(a [3]int) {
	fmt.Println(a)
	a[0] = 123
	fmt.Println(a)
}