package main

import (
	"fmt"
	"time"
)

func main() {
	var tick = time.Tick(500 * time.Millisecond)
	var quit = time.After(3 * time.Second)
FOR:
	for {
		select {
		case <-quit:
			break FOR
		case val := <-tick:
			fmt.Println("tick", val)
		}
	}
}