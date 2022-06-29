package main

import "fmt"

func main() {
	var ch chan int
	ch = make(chan int)
	go func() {
		ch <- 1
	}()
	r := <-ch
	fmt.Println("Received the data from the channel", r)

}

//passing value from sender to receiver   ch<-1(sender)  :=<-ch (receiver)
