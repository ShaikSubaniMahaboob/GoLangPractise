package main

import "fmt"

func main() {
	/*finished(1)
	defer finished(2)
	nums := []int{174, 136, 39, 368, 876, 343, 432, 562, 235}
	largest(nums)

	a := 10
	defer fmt.Println("The deffered one-->", a)
	a = 20
	fmt.Println("normal one-->", a)*/ // defer is used to execute the last in program
	str := "Hello, 世界!"

	fmt.Println()
	for _, v := range str {
		defer fmt.Print(string(v))
	}
}

func finished(i int) {
	fmt.Println("Finished this function-->", i)
}
func largest(nums []int) {
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	fmt.Println("Largest number in", nums, "is", max)
}
