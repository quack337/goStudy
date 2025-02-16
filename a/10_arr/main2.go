package main

import (
	"fmt"
	"unsafe"
)

func main2() { // slice
	var a = []int {1, 2, 3, 4}
	var b = []int {11, 22, 33}
	var c = []int {111, 222}

	fmt.Println(unsafe.Sizeof(a), unsafe.Sizeof(b), unsafe.Sizeof(c))
	fmt.Println(a, b, c)

	a = b
	b[0] = -11
	fmt.Println(a, b, c)

	c = b
	a[0] = -1111
	fmt.Println(a, b, c)

	var d = [][]int {{1, 2}, {3, 4, 5}}

	fmt.Println(unsafe.Sizeof(d))
	fmt.Println(d)

	d[1] = d[0]
	d[0][0] = -11
	fmt.Println(d)

	fmt.Println(unsafe.Sizeof(d))

	var e = []int{1, 2, 3}
	f1(e) // call by ref
	fmt.Println(e)
}

func f1(a []int) {
	fmt.Println(a)
	a[0] = 123
	fmt.Println(a)
}