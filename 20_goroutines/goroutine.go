package main

import (
	"fmt"
	"sync"
)

//	go func(i int) {
//		fmt.Println(i)
//	}(i)
func task(id int, w *sync.WaitGroup) {
	defer w.Done() //defer make the statement run after the function execution
	fmt.Println("Doing task", id)
}
func main() {
	var wg sync.WaitGroup
	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go task(i, &wg)

	}
	wg.Wait()
}
