package main

import (
	"fmt"
	"reflect"
)

func main() {

	var num int8 = 10 // one way to declare a variable
	var num1 = 100    // second way to declare a varfiable
	var (
		num2, num3, isnum = 100.11, "111", false
	) //third way to declare a variable
	numStr := "Number is 500" //shorthand notation of creating a variable by assigning a value
	fmt.Printf("Type of num %T value of num %v\n", num, num)
	fmt.Printf("Type of num1 %T value of num1 %v\n", num1, num1)
	fmt.Printf("Type of num2 %T value of num2 %v\n", num2, num2)
	fmt.Printf("Type of num3 %T value of num3 %v\n", num3, num3)
	fmt.Printf("Type of isnum %T value of isnum %v\n", isnum, isnum)
	fmt.Printf("Type of numStr %T vaalue of numSrt %v\n", numStr, numStr)
	println("\n")
	//swapping
	var a, b, c = 1, 2, 3
	{
		fmt.Println("before swapping a,b,c", a, b, c)
		a, b, c = c, a, b
		fmt.Println("after swapping a,b,c", a, b, c)
	}

	var a1, b1, c1 = 1, 2, 3
	{
		fmt.Println("befaore a1,b1,c1,a1,b1,c1")
		a1, b1, c1 = c1, a1, b1
		fmt.Println("after a1,b1,c1,a1,b1,c1")
	}
	println("\n")
	// type casting

	{
		var long int64
		var small int32 = 500
		long = int64(small)
		fmt.Println("long and small", long, small)
		long1 := 9909

		var small1 int32

		small1 = int32(long1)
		fmt.Println("Long1 and small1", long1, small1)
	}
	println("\n")
	//Empty interface
	{
		var num21 interface{} // this is known as empty interface

		var num11 int

		num21 = 1000

		var num22 = 200

		num11 = num21.(int)

		fmt.Println("Additon of num21+num12", num11+num22)

		fmt.Println("Num is ", num21, " type of num ", reflect.TypeOf(num21))
	}
	println("\n")
	//complex number  [answers will be complex128 or complex64]
	{
		cmplx1 := complex(1000.343, 321.12312)
		fmt.Println("complex number", cmplx1, "type of complex number", reflect.TypeOf(cmplx1))
		var r1, i1 float32 = 1000.343, 321.12312
		cmplx2 := complex(r1, i1)
		fmt.Println("Complex number ", cmplx2, "type of complex number", reflect.TypeOf(cmplx2))
		cmplx3 := complex(1000.345, i1)
		fmt.Println("complex number", cmplx3, "type of complex number", reflect.TypeOf(cmplx3))
		cmplx4 := cmplx3 + complex64(cmplx1)
		fmt.Println("Complex number ", cmplx4, "type of complex number", reflect.TypeOf(cmplx4))
		cmplx5 := cmplx3 - complex64(cmplx1)
		fmt.Println("Complex number ", cmplx5, "type of complex number", reflect.TypeOf(cmplx5))
		cmplx6 := cmplx3 * complex64(cmplx1)
		fmt.Println("Complex number ", cmplx6, "type of complex number", reflect.TypeOf(cmplx6))
		fmt.Println("\t")
		//constants
		const MIN int = 9

		const MAX = 999

		const (
			MAX2 = 999
			MAX3 = 9999
			MIN2 = 99
		)

		var a = (100 * 2) / 2
		fmt.Println("Constant value and normal value", MAX, a)
		const MAX1 = (100 * 9) / MIN
		fmt.Println("Constant value", MAX1)
		const max4 = (MAX2 + MAX3 + MIN2) / 100
		fmt.Println("constant value", max4)
		const max5 = (MAX2 + MAX3) / 2
		fmt.Println("constand value", max5)
	}

}
