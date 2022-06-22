package main

import "fmt"

func main() {
	{
		var a, b, c = 1, 2, 3

		fmt.Println("before swapping a,b,c", a, b, c)
		a, b, c = c, a, b
		fmt.Println("after swapping a,b,c", a, b, c)
	}

}
