package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)

	defer func() {
		close(ch1)
		close(ch2)
	}()

	go func() {
		for {
			time.Sleep(time.Second)
			rand.Seed(time.Now().UnixNano())
			val := rand.Intn(100)
			println("put data to chan1: ", val)
			ch1 <- val
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second)
			rand.Seed(time.Now().UnixNano())
			val := rand.Intn(100)
			println("put data to chan2: ", val)
			ch2 <- val
		}
	}()

	for {
		select {
		case v1 := <-ch1:
			fmt.Printf("recive data from chan1: %d\n", v1)
		case v2 := <-ch2:
			fmt.Printf("recive data from chan2: %d\n", v2)
			//default:
			//	fmt.Printf("not receive data\n")
			//	time.Sleep(time.Microsecond * 100)
		}
	}

	println("main waiting...")
	time.Sleep(time.Minute)
	println("main finish")

}
