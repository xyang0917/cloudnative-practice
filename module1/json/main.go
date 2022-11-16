package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

func main() {
	stu := Student{Name: "yangxin", Age: 30}
	stuStr := object2string(&stu)
	fmt.Printf("%v\n", stuStr)

	stuObj := string2object(stuStr)
	fmt.Printf("%v\n", stuObj)

	object, err := toObject(stuStr)
	if err == nil {
		fmt.Printf("%v\n", object)
	}
}

func string2object(val string) *Student {
	stu := Student{}
	err := json.Unmarshal([]byte(val), &stu)
	if err != nil {
		return nil
	}
	return &stu
}

func object2string(stu *Student) string {
	content, err := json.Marshal(stu)
	if err != nil {
		println("marshal error: ", err)
		return ""
	}

	return string(content)
}

func toObject(humanStr string) (interface{}, error) {
	var obj interface{}
	err := json.Unmarshal([]byte(humanStr), &obj)
	if err != nil {
		return nil, err
	}

	objMap, ok := obj.(map[string]interface{})

	if !ok {
		return nil, errors.New("unmarshal error")
	}

	for k, v := range objMap {
		switch value := v.(type) {
		case string:
			fmt.Printf("type of %s is string, value is %v\n", k, value)
		case interface{}:
			fmt.Printf("type of %s is interface{}, value is %v\n", k, value)
		default:
			fmt.Printf("type of %s is wrong, value is %v\n", k, value)
		}
	}
	return obj, nil
}

type Student struct {
	Name string
	Age  int
}
