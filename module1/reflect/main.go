package main

import (
	"fmt"
	"reflect"
)

func main() {
	/*myMap := make(map[string]string, 10)
	myMap["name"] = "yangxin"
	myMap["sex"] = "man"
	t := reflect.TypeOf(myMap)
	println(t.String())

	v := reflect.ValueOf(myMap)
	for _, key := range v.MapKeys() {
		println(key.String())
	}*/

	// struct reflect
	stu := Student{Name: "yangxin", Age: 30, Sex: "man"}
	stuType := reflect.ValueOf(stu)
	for i := 0; i < stuType.NumField(); i++ {
		field := stuType.Field(i)

		fmt.Printf("Field %d: %v\n", i, field)
	}

	// 只能获取接口中定义的方法
	for i := 0; i < stuType.NumMethod(); i++ {
		method := stuType.Method(i)
		fmt.Printf("Method %d：%v\n", i, method)
	}

	method := stuType.MethodByName("String")
	retValue := method.Call(nil)

	println(retValue[0].String())

}

type Student struct {
	Name string
	Age  int
	Sex  string
}

func (stu Student) String() string {
	return stu.Name + "1"
}

//func (stu Student) setName(name string) {
//	stu.Name = name
//}
