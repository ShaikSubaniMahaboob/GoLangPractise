package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var str rune = 's'
	fmt.Println(int(str))
	var num int = 100
	fmt.Println(string(num))

	nums := fmt.Sprintf("%v", num)
	fmt.Println("value of nums", nums, "type of nums", reflect.TypeOf(nums))
	numstr := strconv.Itoa(num)
	fmt.Println("Value of numstr", numstr, "Type of numstr", reflect.TypeOf(numstr))
	num2, _ := strconv.Atoi("65")
	fmt.Println("num2 using atoi", num2)
}
