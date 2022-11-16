package main

import "fmt"

func main() {
	//var cat Animal = &Cat{Color: "白色"}
	//cat.run()
	//var dog Animal = &Dog{Color: "黑色"}
	//dog.run()

	cat := Cat{Color: "白色"}
	dog := Dog{Color: "黑色"}

	animals := []Animal{&cat, &dog, nil}
	//animals = append(animals, &cat, &dog)
	for _, animal := range animals {
		if animal != nil {
			animal.run()
		}
	}
}

type Animal interface {
	run()
}

type Cat struct {
	Color string
}

func (cat *Cat) run() {
	fmt.Printf("一只%s的猫在奔跑...\n", cat.Color)
}

type Dog struct {
	Color string
}

func (dog *Dog) run() {
	fmt.Printf("一条%s的狗在奔跑...\n", dog.Color)
}
