package main

import (
	"fmt"
	"time"
)

var count int = 0

func main() {
	for i := 1; i <= 1000; i++ {
		go Increment()
		go Decrement()
	}
	time.Sleep(time.Second * 1)
	fmt.Println("The final count", count)

}
func Increment() {
	count++
	fmt.Println("Increment", count)
}
func Decrement() {
	count--
	fmt.Println("Decrement", count)
}
