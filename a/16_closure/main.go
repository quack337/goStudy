package main

import "fmt"

type CounterFn func() int

func createCounter(init int) (CounterFn, CounterFn) {
	return func() int {
			init += 1
			return init
		}, func() int {
			init -= 1
			return init
		}
}

func main() {
	var inc1, dec1 = createCounter(10)
	fmt.Println(inc1())
	fmt.Println(inc1())
	fmt.Println(inc1())
	fmt.Println(inc1())

	var inc2, dec2 = createCounter(5)
	fmt.Println(inc2())
	fmt.Println(inc2())
	fmt.Println(inc2())
	fmt.Println(inc2())

	fmt.Println(dec1())
	fmt.Println(dec1())
	fmt.Println(dec1())

	fmt.Println(dec2())
	fmt.Println(dec2())
	fmt.Println(dec2())
}
