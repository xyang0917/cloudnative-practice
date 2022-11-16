package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

const TimeFormat = "2006-01-02 15:04:05"

func main() {
	ch := make(chan int, 10)
	defer close(ch)

	baseCtx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(baseCtx, time.Second*100)
	defer cancel()

	// 生产者
	go producer(timeoutCtx, ch)

	// 消费者
	go consumer(timeoutCtx, ch)

	timer := time.NewTimer(time.Second * 3)
	defer timer.Stop()
	select {
	case <-timer.C:
		println("main stoping")
		cancel()
		time.Sleep(time.Second * 2)
		println("main stoped")
	}
}

func producer(timeoutCtx context.Context, ch chan<- int) {
	ticker := time.NewTicker(time.Second)

	for range ticker.C {
		select {
		case <-timeoutCtx.Done():
			fmt.Printf("[%s]Stop produce\n", time.Now().Format(TimeFormat))
			return
		default:
			rand.Seed(time.Now().UnixNano())
			n := rand.Intn(1000)
			ch <- n
			fmt.Printf("[%s]produce data: %d\n", time.Now().Format(TimeFormat), n)
		}
	}
}

func consumer(timeoutCtx context.Context, ch <-chan int) {
	ticker := time.NewTicker(time.Second)
	for range ticker.C {
		select {
		case <-timeoutCtx.Done():
			fmt.Printf("[%s]Stop consume\n", time.Now().Format(TimeFormat))
			return
		default:
			for value := range ch {
				fmt.Printf("[%s]consume data: %d\n", time.Now().Format(TimeFormat), value)
			}
		}
	}
}
