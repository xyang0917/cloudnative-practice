package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//StartThread()

	//ThreadCommunication()
	LoopChan()
	//SingleChan()
}

// SingleChan 单向通道
func SingleChan() {
	/*// 只写通道
	var writeOnly chan<- int
	writeOnly = make(chan int, 1)
	writeOnly <- 10
	// error 不能读
	//v := <- writeOnly

	// 只读通道
	var readOnly <-chan int
	readOnly = make(chan int, 1)
	v := <-readOnly
	println(v)
	// error 不能写
	//readOnly <- 10*/

	// 双向通道
	ch := make(chan int, 1)
	doWrite(ch)
	doRead(ch)
}

func doWrite(ch chan<- int) {
	// 双向通道转只写单向通道
	ch <- 10
	println("doWrite finish")
}

func doRead(ch <-chan int) {
	// 双向通道转只读单向通道
	//fmt.Printf("doRead %d\n", <-ch)
	println(<-ch)
	//error 只读通道，不能写
	//ch <- 20
}

func LoopChan() {
	// 创建一个协程通信通道，通信数据为int，通信缓存区大小为2。
	ch := make(chan int, 2)
	go func() {
		// 关闭chan，通知对方不再发送新数据
		defer close(ch)

		for i := 0; i < 20; i++ {
			rand.Seed(time.Now().UnixNano())
			n := rand.Intn(10)
			fmt.Printf("[写]put to chan[%d]: %d\n", i, n)
			ch <- n
		}
	}()

	fmt.Printf("main fun\n")
	for i := range ch {
		fmt.Printf("[读]pull from chan: %d\n", i)
	}

}

func ThreadCommunication() {
	// 默认缓冲区大小为0。当缓存冲区满时，数据的读写操作会阻塞。
	ch := make(chan int)
	go func() {
		println("hello from goroutine")
		// 往管道中写数据
		ch <- 1
	}()

	println("准备从管道中读取数据")

	// 从管理中读数据
	// 如果ch中没有数据，则阻塞当前线程
	data := <-ch

	fmt.Printf("从管道中读取数据:%d\n", data)
}

func StartThread() {
	// go 函数名
	for i := 0; i < 10; i++ {
		go println(i)
	}

	time.Sleep(time.Second)
}
