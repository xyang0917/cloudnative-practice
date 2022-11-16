package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	baseCtx := context.Background()
	ctx := context.WithValue(baseCtx, "a", "b")
	go func(c context.Context) {
		fmt.Printf("Get context value: %s\n", c.Value("a"))
	}(ctx)

	timeoutCtx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	go func(timeout context.Context) {
		ticker := time.NewTicker(time.Second)
		for range ticker.C {
			select {
			case <-timeout.Done():
				println("sub thread finish.")
				return
			default:
				fmt.Printf("attr value: %s\n", timeout.Value("a"))
			}
		}
	}(timeoutCtx)

	time.Sleep(time.Second * 6)
	println("main finish")

}
