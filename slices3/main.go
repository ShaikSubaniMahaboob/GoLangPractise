package main

import (
	"fmt"
)

func main() {
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("slice1", slice1)
	slice2 := slice1
	slice2[0] = 101
	slice2[1] = 102
	slice2[2] = 103
	fmt.Println("slice2", slice2)

	//clone a slice (shallow copy)
	slice3 := make([]int, len(slice1))
	for index, val := range slice1 {
		slice3[index] = val
	}
	slice3[0] = 1011
	slice3[1] = 1022
	slice3[2] = 1033
	fmt.Println("slice1", slice1)
	fmt.Println("slice3", slice3)
	//clone a slice by using copy function (deep function)
	slice4 := make([]int, len(slice1))
	copy(slice4, slice1)
	fmt.Println("slice4", slice4)
	slice4[0] = 1011
	slice4[1] = 1022
	slice4[2] = 1033
	fmt.Println("slice1", slice1)
	fmt.Println("slice4", slice4)

	slice5 := slice1[5:]
	fmt.Println("slice5", slice5)

	slice6 := slice1[:5]
	fmt.Println("slice6", slice6)

	slice6[0] = 11111
	fmt.Println("slice6", slice6, "slice1", slice1)

	slice7 := slice1[3:7]
	fmt.Println("slice7", slice7)

	//convert array to slice (deep slice)
	arr := [5]int{1, 2, 3, 4, 5}
	//fmt.Println(arr)
	slice9 := arr[:] //this is refernece type
	slice10 := arr   // this is value type
	slice9[0] = 101
	slice10[0] = 1111
	fmt.Println("Array", arr)
	fmt.Println("slice9", slice9)
	fmt.Println("Slice10", slice10)
	//remove an eliment from slice
	slice9 = remove(slice9, 0)
	fmt.Println("slice9", slice9)

	slice9 = remove(slice9, 1)
	fmt.Println("slice9", slice9)

	slice7 = remove(slice7, 0)
	fmt.Println("slice7", slice7)

}

func remove(slice []int, i int) []int {
	slice = append(slice[:i], slice[i+1:]...)
	return slice
}
