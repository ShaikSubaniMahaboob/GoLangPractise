package main

import (
	"fmt"
	"reflect"
)

func main() {
	cmpl1 := complex(100.10, 200.22)
	fmt.Println("complex number", cmpl1, "type of complex number", reflect.TypeOf(cmpl1))
}
