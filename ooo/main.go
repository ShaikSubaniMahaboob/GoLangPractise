package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go greet(&wg)

	}

	fmt.Println("Main method")
	wg.Wait()

}
func greet(wg *sync.WaitGroup) {
	fmt.Println("hello world from greet function")
	wg.Done()
}
