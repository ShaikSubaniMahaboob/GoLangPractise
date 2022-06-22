package main

import "fmt"

func main() {
	var long int64
	var small int32 = 500
	long = int64(small)
	fmt.Println("long and small", long, small)
	long1 := 9909

	var small1 int32

	small1 = int32(long1)
	fmt.Println("Long1 and small1", long1, small1)
}
