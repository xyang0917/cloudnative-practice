package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	condAll()
}

func condAll() {
	mutex := sync.Mutex{}
	cond := sync.NewCond(&mutex)
	for i := 0; i < 10; i++ {
		go func(i int) {
			cond.L.Lock()
			defer cond.L.Unlock()
			cond.Wait()
			fmt.Println("Wait", i)
			time.Sleep(time.Second)
		}(i)
	}

	time.Sleep(time.Second)
	cond.Broadcast()

	time.Sleep(time.Second * 15)

	fmt.Println("main exit")
}

func condSingle() {
	mutex := sync.Mutex{}
	cond := sync.NewCond(&mutex)
	cond.L.Lock()

	go func() {
		fmt.Printf("Sub thread executing...\n")
		time.Sleep(3 * time.Second)
		cond.Signal()
	}()

	fmt.Printf("Main thread wait begin\n")

	cond.Wait()

	fmt.Printf("Main thread wait end\n")
}
