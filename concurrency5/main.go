package main

import "fmt"

var done chan bool

func main() {
	done = make(chan bool)
	go greet()
	r := <-done
	fmt.Println("Hello Folks!,This is main method", r)

}
func greet() {
	fmt.Println("Hello Folks!,This is greet method")
	done <- false
}
