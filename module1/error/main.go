package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	var x, y = 10, 0
	result, err := div(x, y)
	if err == nil {
		fmt.Printf("%d / %d = %d\n", x, y, result)
	} else {
		println(err.Error())
	}

	defer0()
}

func div(x, y int) (int, error) {
	if y == 0 {
		//err := fmt.Errorf("除数不能为0")
		err := errors.New("除数不能为0")
		return -1, err
	}

	result := x / y

	return result, nil
}

func defer0() {
	defer func() {
		println("defer0函数执行结束回调")
		if err := recover(); err != nil {
			fmt.Printf("defer0函数执行错误：%s\n", err)
		} else {
			println("函数中未引发异常")
		}
	}()

	println("defer0函数开始执行")
	time.Sleep(3000)
	println("defer0执行结束")

	//panic("模拟执行异常")

}
