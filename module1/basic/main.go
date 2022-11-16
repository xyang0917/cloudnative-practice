package main

import (
	"fmt"
	"strings"
)

func main() {
	//loop()

	// 常量
	/*const HOST string = "localhost"
	// 变量
	var ip string
	var port int
	ip = "127.0.0.1"
	port = 8080
	fmt.Printf("%s:%d\n", ip, port)

	var c, java, php bool
	println(c, java, php)

	// 定义并赋值
	name := "yangxin"
	println(name)*/

	/*f := 200
	//i := int16(f)
	i := f
	println(i)*/

	//map0()

	/*stu := Student{stuNo: "10001", name: "yangxin", age: 30}
	//stu.age = 40
	//s := toString(&stu)
	//println(s)
	stuType := reflect.TypeOf(stu)
	name, b := stuType.FieldByName("name")
	fmt.Printf("%v, %v\n", name, b)

	tag := name.Tag.Get("json")
	fmt.Printf("tag = %s\n", tag)

	var myName MyString = "yangxin"

	println("myName", myName)*/

	arr()
}

type MyString string

type Student struct {
	stuNo string `json:"stuNo"`
	name  string `json:"name22"`
	age   int    `xml:"age"`
}

func toString(stu *Student) string {
	return fmt.Sprintf("stuNo=%s, name=%s, age=%d\n", stu.stuNo, stu.name, stu.age)
}

func init() {
	println("main init")
}

func map0() {
	var user = make(map[string]string, 10)
	user["name"] = "yangxin"
	user["age"] = "30"
	/*fmt.Printf("%v\n", user)
	fmt.Println(user["name"])

	myFunMap := map[string]func() int{
		"funcA": func() int {
			return 100
		},
		"funcB": func() int {
			return 200
		},
	}

	funcA := myFunMap["funcA"]
	value := funcA()
	println(value)*/

	value, exists := user["name2"]
	println(value, exists)

	for key, value := range user {
		fmt.Printf("%s = %s\n", key, value)
	}
}

func slice() {
	//myArray := [5]int{1, 2, 3, 4, 5}
	//fmt.Printf("array = %v\n", myArray)
	//mySlice := myArray[:3]
	//fmt.Printf("slice = %v\n", mySlice)

	mySlice := make([]int, 10, 20)
	fmt.Printf("%d - %v\n", len(mySlice), mySlice)
	for i, _ := range mySlice {
		mySlice[i] += 2
	}
	println("change after")
	for i, v := range mySlice {
		fmt.Printf("%d = %d\n", i, v)
	}

	//mySlice2 := new([]int)
	//fmt.Printf("%v\n", mySlice2)

}

func arr() {
	/*var food [10]string
	food[0] = "宫保鸡丁"

	num := [3]int{1, 2, 3}
	for i, v := range num {
		fmt.Printf("%d = %d\n", i, v)
	}

	strs := [3]string{"yangxin", "yaokun", "yangyu"}
	for i, v := range strs {
		fmt.Printf("%d = %s\n", i, v)
	}

	println(strs[0])*/

	arrs := [10]string{"I", "am", "stupid", "and", "weak"}
	fmt.Printf("修改前：%v\n", arrs)
	for i, _ := range arrs {
		value := arrs[i]
		if strings.Compare(value, "stupid") == 0 {
			arrs[i] = "smart"
		} else if strings.Compare(value, "weak") == 0 {
			arrs[i] = "strong"
		}
	}

	fmt.Printf("修改后：%v\n", arrs)
}

func loop() {
	/*for i := 0; i < 100; i++ {
		fmt.Printf("%d\n", i)
	}*/

	//i := 0
	/*for i < 1000 {
		i++
	}
	fmt.Printf("loop in %d\n", i)*/

	// 死循环
	/*for {
		if i > 1000 {
			break
		}
		i++
	}*/

	// 遍历字符串
	var str = "hello world"
	for index, char := range str {
		fmt.Printf("[%d] = %c\n", index, char)
	}

	// 遍历数组
	//var arr = [0, 1,2,3]
}

func switch0() {
	var x = 90
	fmt.Printf("x=%d\n", x)
	switch x {
	case 90, 100:
		fmt.Println("优秀")
		fallthrough
	case 70, 80:
		fmt.Println("良好")
	case 60:
		fmt.Println("及格")
	default:
		fmt.Println("不及格")
	}
}

func ifelse() {
	var x = 30
	if x == 10 {
		fmt.Printf("this is %d\n", x)
	} else if x == 20 {
		fmt.Printf("this is 20")
	} else {
		fmt.Println("this is other")
	}

	if v := x - 10; v < 100 {
		fmt.Printf("v少于100")
	}
}
