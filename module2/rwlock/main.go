package main

import (
	"fmt"
	"sync"
	"time"
)

var data int
var lock = sync.RWMutex{}

func main() {
	go read()
	go read()
	go write()
	go read()

	time.Sleep(time.Second * 5)
}

func read() {
	lock.RLock()
	defer lock.RUnlock()
	time.Sleep(time.Second * 1)
	fmt.Println("Read data is:", data)
}

func write() {
	lock.Lock()
	defer lock.Unlock()
	data = 100
	fmt.Printf("Write data is:%d\n", data)
}
