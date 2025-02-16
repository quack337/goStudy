package main

import (
	"fmt"
)

func main2() {

	for i := 0; i < 10; i++ {
		func() {
			defer func() { if e := recover(); e != nil { fmt.Println(e) } }()
			f2(i)
		}()
	} 
}

func f2(i int) {
	if i % 2 == 1 { panic(fmt.Sprintf("odd number %d", i)) }
	fmt.Printf("%d\n", i)
}
