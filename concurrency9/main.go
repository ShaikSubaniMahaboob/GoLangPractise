package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		for i := 1; i <= 1000; i++ {
			ch <- 1
		}
		close(ch)
	}()
	for v := range ch {
		fmt.Println(v)
	}

}
