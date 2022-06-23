package main

import (
	"fmt"
	"math/rand"
)

func valuechange(arr []int, index uint, val int) {
	if index < uint(len(arr)) {
		arr[index] = val
	}
	fmt.Println("Array Inside function", arr)
}
func main() {
	arr := []int{22, 33, 44}
	fmt.Println("Array", arr)
	valuechange(arr, 2, 99)
	fmt.Println("after changes", arr)
	arr = []int{88, 98, 978, 78}
	fmt.Println(arr)
	//Slice
	var slice []int //declaring slice
	fmt.Println("Type of slice", slice)
	fmt.Println("Length of Slice", len(slice), "\n Capacity of Slice", cap(slice))
	if slice == nil {
		fmt.Println("if the slice is nil,it to be Instantiated")
	}
	slice = make([]int, 10, 10)
	if slice == nil {
		fmt.Println("if the slice is nil,it to be Instantiated")
	} else {
		fmt.Println("Len", len(slice), "\ncap", cap(slice))
		fmt.Println("slice", slice)
	}
	for i := 0; i < len(slice); i++ {
		slice[i] = rand.Intn(900)
	}
	fmt.Println("slice", slice)

	slice = append(slice, 1000)
	slice = append(slice, 1001)
	for i := 0; i < len(slice); i++ {
		fmt.Println("Index", i, "Value", slice[i])
	}
	fmt.Println("Length", len(slice), "Capacity", cap(slice))
	fmt.Println("slice", slice)

	slic2 := make([]string, 10) //shorthand notation
	slic2[0] = "Bangalroe"
	slic2[1] = "Hyderabad"
	slic2[3] = "Guntur"
	slic2[4] = "kakinada"
	fmt.Println("slice2", slic2)
	slic2 = append(slic2, "king")
	fmt.Println("slice2", slic2)

	var slice3 []string

	slice3 = append(slice3, "hello")
	fmt.Println(" slice3", slice3)

}
