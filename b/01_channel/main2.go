package main

import (
	"fmt"
	"sync"
)

var wg2 sync.WaitGroup

func woker2(data chan int) {
	for val := range data {
		fmt.Println(val)
	}	
	fmt.Println("done")
	wg2.Done()
}

func main2() {
	var data = make(chan int, 5)
	wg2.Add(1)
	go woker2(data)
	for i := 100; i <= 105; i++ {
		data <- i
	}
	close(data)
	wg2.Wait()
	fmt.Println("quit")
	close(data)
}