package main

import (
	"fmt"
	"math/rand"
)

func main() {
	day := 73
	switch day {
	case 1:
		fmt.Println("Sunday")

	case 2:
		fmt.Println("Monday")

	case 3:
		fmt.Println("Tuesday")
	case 4:
		fmt.Println("Wednesday")

	case 5:
		fmt.Println("Thursday")

	case 6:
		fmt.Println("Friday")
	case 7:
		fmt.Println("Saturday")
	default:
		fmt.Println("Noday")
	}

	number := 6
	switch {
	case number%8 == 0:
		fmt.Println(number, "is divisible of 8")
		fallthrough
	case number%6 == 0:
		fmt.Println(number, "is divisible by 6")
		fallthrough
	case number%4 == 0:
		fmt.Println(number, "is divisible by 4")

	}
	number = 55
	switch {
	case number <= 50 && number >= 0:
		fmt.Println(number, "is between 0 and 50")
	case number >= 50 && number <= 100:
		fmt.Println(number, "is between 50 and 100")
	case number >= 100:
		fmt.Println(number, "is greater than 100")
	}

	char := 'e'
	switch char {
	case 'a', 'e', 'i', 'o', 'u':
		fmt.Println(string(char), "is vowel")
	}
	//number =100
	for i := 0; i <= 10; i++ {
		switch number := rand.Intn(100); number % 2 {
		case 1:
			fmt.Println(number, "print only even number")
		case 2:
			fmt.Println(number, "print only odd number")

		}

	}
}
