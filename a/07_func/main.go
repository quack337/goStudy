package main

import "fmt"

func sum1(a []int) int {
	var result = 0
	for _, value := range a {
		result += value
	}
	return result
}

func sum2(a []int) (result int) {
	for i := 0; i < len(a); i++ {
		result += a[i]
	}
	return result
}

func main() {
	var a = []int {1, 2, 3, 4, 5}
	fmt.Println(sum1(a))
	fmt.Println(sum2(a))
}