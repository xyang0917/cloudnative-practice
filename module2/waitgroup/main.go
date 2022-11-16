package main

import (
	"fmt"
	"sync"
)

func main() {
	group := sync.WaitGroup{}
	group.Add(10)

	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Printf("Done %d\n", i)
			group.Done()
		}(i)
	}

	group.Wait()
	fmt.Println("Wait end")
}
