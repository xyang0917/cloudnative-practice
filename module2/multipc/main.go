package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var queue []int

	mutex := sync.RWMutex{}
	cond := sync.NewCond(&mutex)

	// 生产者
	for i := 0; i < 3; i++ {
		go producer(cond, &queue, i)
	}

	// 消费者
	for i := 0; i < 5; i++ {
		go consumer(cond, &queue, i)
	}

	time.Sleep(time.Second * 60)
	println("main exit", len(queue))

}

func producer(cond *sync.Cond, queue *[]int, i int) {
	for {
		func(i int) {
			cond.L.Lock()
			defer cond.L.Unlock()

			rand.Seed(time.Now().UnixMilli())
			time.Sleep(time.Second)
			data := rand.Intn(1000)
			*queue = append(*queue, data)
			log.Printf("[%d]Produce data: %d，堆积长度:%d\n", i, data, len(*queue))

			cond.Broadcast()
		}(i)
	}
}

func consumer(cond *sync.Cond, queue *[]int, i int) {
	for {
		func(i int) {
			cond.L.Lock()
			defer cond.L.Unlock()

			for len(*queue) == 0 {
				//log.Printf("[%d]Not data can consume, waiting...", i)
				cond.Wait()
			}
			data := (*queue)[0]
			*queue = (*queue)[1:]
			log.Printf("[%d]Consume data: %d\n", i, data)
		}(i)
	}
}
