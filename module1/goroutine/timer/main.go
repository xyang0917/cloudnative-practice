package main

import (
	"fmt"
	"time"
)

func main() {

	//go timeout()

	go ticker0()

	println("main begin")
	time.Sleep(time.Second * 10)
	println("main finish")

}

// 定时器
func ticker0() {
	// 每2秒执行一次
	ticker := time.NewTicker(time.Second * 2)
	defer ticker.Stop()

	for {
		select {
		case v := <-ticker.C:
			fmt.Printf("ticker in : %v\n", v)
		}
	}

}

// 超时器，只执行一次
func timeout() {
	// 1秒后执行
	timer := time.NewTimer(time.Second * 1)
	defer timer.Stop()

	for {
		select {
		case v := <-timer.C:
			fmt.Printf("timeout in : %s\n", v.String())
			break
		}
	}

}
