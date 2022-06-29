package main

import "fmt"

func main() {
	data := make(chan string, 2) // this is called buffered channel
	data <- "hello"
	data <- "How"
	// mention size above 2 only
	fmt.Println(<-data)
	data <- "Are"
	fmt.Println(<-data)
	data <- "you"
	fmt.Println(<-data)
	fmt.Println(<-data)

}
