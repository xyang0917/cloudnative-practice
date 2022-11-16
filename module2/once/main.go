package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	// 一个Once实例，只会执行一次
	once := sync.Once{}
	onceFun := func() {
		fmt.Printf("Only exec once")
	}

	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceFun)
		}()
	}

	time.Sleep(time.Second * 3)

}
