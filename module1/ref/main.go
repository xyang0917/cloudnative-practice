package main

import "fmt"

func main() {
	stu := Student{Name: "zhangsan", Age: 20}
	fmt.Printf("%v\n", stu)

	changeName(&stu)
	fmt.Printf("%v\n", stu)

	changeAge(stu)
	fmt.Printf("%v\n", stu)

	var x, y, result int = 10, 20, 0
	calc(x, y, &result)
	fmt.Printf("%d + %d = %d", x, y, result)

}

type Student struct {
	Name string
	Age  int
}

func calc(x, y int, result *int) {
	*result = x + y
}

// 传引用
func changeName(stu *Student) {
	stu.Name = "yangxin01"
}

// 传值
func changeAge(stu Student) {
	stu.Age = 30
}
