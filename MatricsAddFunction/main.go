package main

import "fmt"

func main() {
	var rows, columns, i, j int

	var addmat1 [10][10]int
	var addmat2 [10][10]int
	var additionmat [10][10]int

	fmt.Print("Enter the Addition Matrix Rows and Columns = ")
	fmt.Scan(&rows, &columns)

	fmt.Print("Enter the First Matrix Items to Add = ")
	for i = 0; i < rows; i++ {
		for j = 0; j < columns; j++ {
			fmt.Scan(&addmat1[i][j])
		}
	}

	fmt.Print("Enter the Second Matrix Items to Add = ")
	for i = 0; i < rows; i++ {
		for j = 0; j < columns; j++ {
			fmt.Scan(&addmat2[i][j])
		}
	}

	for i = 0; i < rows; i++ {
		for j = 0; j < columns; j++ {
			additionmat[i][j] = addmat1[i][j] + addmat2[i][j]
		}
	}
	fmt.Println("The Sum of Two Matrices = ")
	for i = 0; i < rows; i++ {
		for j = 0; j < columns; j++ {
			fmt.Print(additionmat[i][j], "\t")
		}
		fmt.Println()
	}

}
