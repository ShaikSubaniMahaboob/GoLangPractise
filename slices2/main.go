package main

import "fmt"

func main() {
	var slice1 = make([]int, 10)
	slice1[0] = 1
	slice1[1] = 2
	slice1[2] = 3
	slice1[3] = 4
	slice1 = append(slice1, 3, 4, 5)
	fmt.Println("Slice1", slice1)

	var slice2 []string
	slice2 = append(slice2, "hello")
	slice2 = append(slice2, "how", "are", "you")
	fmt.Println(slice2)
	// slices contain AND RETURNS TWO parametr index ,value
	//for i := 0; i <= 10; i++ {

	for i, val := range slice2 {
		fmt.Println("index", i, "value", val)
	}
	//}
	//need only index
	for i1, _ := range slice2 { //blank identifier
		fmt.Print("\nindex", i1)
	}
	fmt.Println("\n")
	//if you need only value
	for _, val1 := range slice2 { //blank identifier
		fmt.Println("", val1)
	}

	slice3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("slice3", slice3)
	changeVal(slice3, 2, 300, 400, 500, 600, 700, 800, 900, 1000, 1100, 1200, 1300)
	fmt.Println("Slice3", slice3)
	vslice := []int{300, 400, 500, 606, 707, 808, 909, 1001, 1100, 1200, 1300}
	changeVal(slice3, 2, vslice...)
	fmt.Println("Slice3", vslice)
}
func changeVal(slice []int, sindex int, vals ...int) {
	if slice != nil {
		if sindex < len(slice) && len(vals) >= 1 {
			for i, j := sindex, 0; i < len(slice) && j < len(vals); i, j = i+1, j+1 {
				slice[i] = vals[j]
			}
		}
	}
}
