package main

import (
	"fmt"
)

func main() {
	greet()
	greetby("subani")
	greetwith("hello", "subani")

	//return type i have used method name in short hand declaration
	a, p := AreaAndPeriOfRect(100.45, 108.74)
	fmt.Println("area", a, "perimeter", p)

	a1, _ := AreaAndPeriOfRect(100.45, 108.74) // blank identifier
	fmt.Println("area", a1)

	_, p1 := AreaAndPeriOfRect(100.45, 108.740) // blank identifier
	fmt.Println("perimeter", p1)

}
func greet() { //without parameters
	fmt.Println("Hello method")
}
func greetby(name string) {
	fmt.Println("hello", name)
}
func greetwith(greet string, name string) {
	fmt.Println(greet, name)
}

//Return type

func AreaOfRect(l, b float32) float32 {
	return l * b
}
func PerimeterOfRect(l float32, b float32) float32 {
	return 2 * (l + b)
}
func AreaAndPeriOfRect(l, b float32) (area float32, perimeter float32) {
	area = l * b
	perimeter = 2 * (l + b)
	return area, perimeter
}
