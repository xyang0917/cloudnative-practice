// 函数
package main

import (
	"errors"
	"flag"
	"fmt"
)

var name = flag.String("name", "jake", "学生姓名")
var age = flag.Int64("age", 0, "学生年龄")

func main() {
	/*args := os.Args
	fmt.Printf("%v\n", args)

	flag.Parse()
	fmt.Printf("name=%s, age=%d, other=%v\n", *name, *age, flag.Args())*/

	//append0()

	//result, err := inc(0, 10)
	//if err == nil {
	//	fmt.Printf("result=%d\n", result)
	//} else {
	//	println(err.Error())
	//}

	//value := sum(1, 2, 3, 4, 5)
	//println(value)

	/*url := "https://www.baidu.com"
	res := request(url, func(s string) string {
		return s + "--aaaaaa"
	})
	println(res)*/

	defer func() {
		println("exit")
	}()

	i := func(x, y int) int {
		println(x, y)
		return x + y
	}(10, 20)
	println(i)

	x := func() { println("this is closure fun") }
	x()

}

func append0() {
	myArray := []string{"0"}
	values := append(myArray, "a", "b", "c")
	fmt.Printf("%v\n", values)

	println(cap(myArray), len(myArray))
}

func inc(x, y int) (int, error) {
	if x != 0 && y != 0 {
		return x + y, nil
	}

	return 0, errors.New("参数错误")
}

func sum(n ...int) int {
	result := 0
	for i := range n {
		result += n[i]
	}
	return result
}

func request(url string, handle func(string) string) string {
	println("request url: ", url)
	data := "{'result': 'ok'}"
	response := handle(data)
	return response
}
