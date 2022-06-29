package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)
	go greet(done)
	data := make(chan string)
	go publisher(data)
	go subscriber(data)

	<-done
	time.Sleep(time.Millisecond * 10)
	fmt.Println("This is main method")

}
func greet(ok chan<- bool) { // sending data in parameters
	fmt.Println("this is greet route method")
	ok <- true // sending data
}
func publisher(send chan<- string) {
	fmt.Println("this is publisher method")
	send <- "name"

}
func subscriber(reciver chan<- string) {
	fmt.Println("this is subscribrt method")
	data := reciver
	fmt.Println("subscriber", data)
}
