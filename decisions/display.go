package main

import "math/rand"

func main() {
	//age := 10  short hand notation
	var age uint = 33
	if age >= 18 {
		println("majar")
	} else {
		println("minor")
	}
	var gender string = "female"
	if age >= 18 && gender == "female" {
		println("she is eligible for marraige")
	} else if age >= 21 && gender == "male" {
		println("He is eligible for marraige")
	} else if gender != "male" && gender != "female" {
		println("No idea")
	} else {
		println("Otherwise not eligible for marriage")
	}

	//declaring and using conditions

	if num1, num2 := rand.Intn(1001), rand.Intn(2000); (num1+num2)%2 == 0 {
		println((num1 + num2), "even number")

	} else {
		println((num1 + num2), "odd  number")
	}
}
