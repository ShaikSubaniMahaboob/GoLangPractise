package main

import (
	"fmt"
)

func main() {
	var ptr *int
	num := 100
	fmt.Println("num", num, "ptr", ptr)
	ptr = &num
	fmt.Println("Value of num", num, "value of num through pointer", *ptr, "address of num", &num, "address os num through pointer", ptr)
	fmt.Println("Address of pointer", &ptr)
	fmt.Println("Value of num through its address", *(&num))
	var ptrOfptr **int

	ptrOfptr = &ptr

	fmt.Println("Pointer of pointer", ptrOfptr)
	fmt.Println("Value of Pointer of pointer ", **ptrOfptr)

	//
	num1 := 100
	fmt.Println("Value before changing", num1)
	changeVal(num1, 200) //all normal values are call by value so that they are passed as values .. a separate copy is maintained
	//fmt.Println("before  value changing", num1)
	// using pointer
	fmt.Println("Value before changing", num1)

	changeValptr(&num1, 300)
	fmt.Println("after before changing", num1)

	ptr3 := new(int)
	fmt.Println("The value of pointer ptr3", *ptr3)
}
func changeVal(variable, value int) {
	variable = value
	fmt.Println("Value after changing inside function", variable)
}
func changeValptr(variable *int, value int) {
	*variable = value
	fmt.Println("Value after changing inside function", *variable)
}
