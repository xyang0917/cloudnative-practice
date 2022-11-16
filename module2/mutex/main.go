package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var data int
var lock = sync.Mutex{}

func main() {
	go read()
	go write()
	go read()

	println("cpu num", runtime.NumCPU())

	time.Sleep(time.Second * 3)
}

func read() {
	lock.Lock()
	defer lock.Unlock()
	time.Sleep(time.Second)
	fmt.Printf("Read data in: %d\n", data)
}

func write() {
	lock.Lock()
	defer lock.Unlock()
	data = 100
	fmt.Println("Write data in : ", data)
}
