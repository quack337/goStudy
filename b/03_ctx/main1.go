package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg1 sync.WaitGroup

func worker1(ctx context.Context) {
	var tick = time.Tick(1 * time.Second)
	for {
		select {
		case <- tick:
			fmt.Println("tick")
		case <- ctx.Done():
			wg1.Done()
			return
		}
	}
}

func main1() {
	var ctx, cancel = context.WithCancel(context.Background())
	wg1.Add(1)
	go worker1(ctx)
	time.Sleep(5 * time.Second)
	cancel()
	wg1.Wait()
}
