package main

import "fmt"

func main4() { // slice
	var a = [...]int {0, 1, 2, 3, 4, 5, 6, 7, 8, 9} // array
	fmt.Println(len(a), cap(a), a);

	var b = a[1:5]
	fmt.Println(len(b), cap(b), b);

	a[2] = 22
	b[2] = 222
	fmt.Println(len(a), cap(a), a);
	fmt.Println(len(b), cap(b), b);
}