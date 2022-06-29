package main

import (
	"fmt"
	"time"
)

var count int = 0
var ch chan int

func main() {
	ch = make(chan int)

	go func() {
		ch <- count
	}()
	func() {
		for i := 0; i < 1000; i++ {
			go Increment()
			go Decrement()
		}
	}()
	time.Sleep(time.Microsecond * 20)
	fmt.Println("Final Count<-", count)
}

func Increment() {
	count = <-ch
	count++
	ch <- count
	fmt.Println("Increment<-", count)

}
func Decrement() {
	count = <-ch
	count--
	ch <- count
	fmt.Println("Decrement<-", count)

}
