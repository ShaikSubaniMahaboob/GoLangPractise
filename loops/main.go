//for
//for-range
//go-to

package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i, " ")
	}
	fmt.Println("\t")
	for i := 2; i <= 10; i = i + 2 {
		fmt.Println(i)
	}
	fmt.Println("\t")
	fmt.Println()
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Print(i, " ")
	}
	fmt.Println("\t")
	for i := 1; i <= 20; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println(i, " ")
	}
	for i := 0; ; {
		if i >= 10 {
			break
		}
		fmt.Print(i, " ")
		i++
	}
	fmt.Print("\n")

	i := 0
	for {
		if i >= 10 {
			break
		}
		fmt.Print(i, " \n")
		i++
	}
	fmt.Println()
	// init two variables and perform loop
	for i, j := 0, 0; i <= 100 && j <= 100; i, j = i+1, j+5 {
		fmt.Println("value of i", i, "value of j", j)

	}
	// using go-to loop-statement

	/*var k = 10
	fmt.Println("k")
	k++
	if k <= 10 {
		goto loop
	}
	fmt.Println("\ngoto loop")*/

}
