package main

import (
	"fmt"
	"reflect"
)

func main() {
	var arr [3]int
	fmt.Println(arr)

	var arr1 [3]int
	arr1[0] = 101
	arr1[1] = 102
	arr1[2] = 103
	//arr1[3] = 105
	fmt.Println(arr1)

	arr2 := [3]int{44, 55, 66}
	fmt.Println(arr2)
	fmt.Println("type of an array", reflect.TypeOf(arr2))
	//its indicate that length of an array
	for i := 0; i < len(arr2); i++ {
		fmt.Print(" ", arr2[i])
	}
	fmt.Println("\nLength of an array", len(arr2))
	fmt.Println("type of aan array", reflect.TypeOf(arr2))
	for i := 1; i < len(arr2); i++ {
		arr2[i] = 100 + i
	}
	fmt.Println("values of arr2 after reassining", arr2)
	fmt.Print("type of an array", reflect.TypeOf(arr2))
	var arr3 = [5]int{111, 222, 333, 444, 555}
	fmt.Println(arr3)
	fmt.Println("type of an array", reflect.TypeOf(arr3))
	arr = arr1 //valid type of arr is [3]int and type of arr1 is [3]int same
	fmt.Println("values of arr", arr, reflect.TypeOf(arr), arr)
	fmt.Println("types of arr1", arr1)
	arrCopy(arr3, arr1)
	//value change function
	arr4 := [4]int{90, 91, 92, 93}
	fmt.Println("Array", arr4)
	valuechange(arr4, 2, 300)
	fmt.Println("after changes", arr4)
}
func arrCopy(src [5]int, dst [3]int) {
	if len(src) <= len(dst) {
		for i := 0; i < len(src); i++ {
			dst[i] = src[i]
		}
	}
}

//change value function
func valuechange(arr4 [4]int, index uint, val int) {
	if index < uint(len(arr4)) {
		arr4[index] = val
	}
	fmt.Println("Array inside function", arr4)
}
