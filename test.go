package main

import "fmt"

func test(matrix [][]int) {

	rows := len(matrix)    // len(matrix) returns the number of rows
	cols := len(matrix[0]) // Note that matrix[0] gives the first row and len(matrix[0])
	//gives number of columns in first row.

	fmt.Printf("Rows: %v Columns %v", rows, cols)
}
