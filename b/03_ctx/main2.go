package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg2 sync.WaitGroup

func worker2(ctx context.Context) {
	var msg = ctx.Value("msg")

	var tick = time.Tick(1 * time.Second)
	for {
		select {
		case <- tick:
			fmt.Println(msg)
		case <- ctx.Done():
			wg2.Done()
			return
		}
	}
}

func main2() {
	var ctx, _ = context.WithTimeout(context.Background(), 5 * time.Second)
	ctx = context.WithValue(ctx, "msg", "hello")
	wg2.Add(1)
	go worker2(ctx)
	wg2.Wait()
}
