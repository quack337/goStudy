package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64 = 0.1
	var b float64 = 0.2 
	var c float64 = 0.3 

	fmt.Printf("%.20f %.20f %.20f %.20f\n", a, b, a+b, c)
	fmt.Println(a+b == c)
	fmt.Println(math.Nextafter(a+b, c) == c)
}
