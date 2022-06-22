package main

import "fmt"

func main() {
	// maltriplication matrix
	//Take number of rows and columns as input.
	//Apply nested for loop to take matrix elements as input.
	//Apply multiplication on matrix elements.
	//Print the resultant matrix.
	var rows, columns, i, j int

	var multimat1 [10][10]int
	var multimat2 [10][10]int
	var multiplicationnmat [10][10]int

	fmt.Print("Enter the Multiplication Matrix Rows and Columns = ")
	fmt.Scan(&rows, &columns)

	fmt.Print("Enter the First Matrix Items to Multiplication = ")
	for i = 0; i < rows; i++ {
		for j = 0; j < columns; j++ {
			fmt.Scan(&multimat1[i][j])
		}
	}

	fmt.Print("Enter the Second Matrix Items to Multiplication = ")
	for i = 0; i < rows; i++ {
		for j = 0; j < columns; j++ {
			fmt.Scan(&multimat2[i][j])
		}
	}

	for i = 0; i < rows; i++ {
		for j = 0; j < columns; j++ {
			multiplicationnmat[i][j] = multimat1[i][j] * multimat2[i][j]
		}
	}
	fmt.Println("The Go Result of Matrix Multiplication = ")
	for i = 0; i < rows; i++ {
		for j = 0; j < columns; j++ {
			fmt.Print(multiplicationnmat[i][j], "\t")
		}
		fmt.Println()
	}
}
