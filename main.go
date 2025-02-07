package main

import "fmt"

func main() {

	var n int
	fmt.Println("Enter matrix size: ")
	fmt.Scan(&n)

	matrix := make([][]int, n) // Create n rows, columns not allocated
	for i := range matrix {
		matrix[i] = make([]int, n) //Iterate over n rows and create n columns per row
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Scanf("%d", &matrix[i][j])
		}
	}

	println("Generated Matrix: ")
	for _, row := range matrix {
		fmt.Println(row)
	}

}
