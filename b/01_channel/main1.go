package main

import (
	"fmt"
	"sync"
	"time"
)

var wg1 sync.WaitGroup

func woker1(data chan int, stop chan bool) {
	defer wg1.Done()
	for {
		time.Sleep(100 * time.Millisecond)
		select {
		case <- stop: // 먼저 적은 것이 우선 순위가 높다
			fmt.Println("receive stop")
			return
		case val := <- data:
			fmt.Println(val)
		}
	}
}

func main1() {
	var data = make(chan int, 5) // 채널 크기 5
	var stop = make(chan bool, 1) // 채널 크기가 0 이면 기다리게 되겟군
	wg1.Add(1)
	go woker1(data, stop)
	for i := 100; i <= 110; i++ {
		data <- i
	}
	stop <- true // 106, 107, 108, 109, 110 이렇게 채우자 마자 stop을 채움
	             // 따라서 105까지는 확실히 처리되지만, 나머지는 ...
				 // 채널 크기가 0 이면 받아갈 때까지 여기서 기다리게 됨				 
	fmt.Println("send stop")				 
	wg1.Wait()
	close(data)
	close(stop)
}