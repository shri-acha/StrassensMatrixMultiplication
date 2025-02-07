package main

import "fmt"

func main() {

	var n int
	fmt.Println("Enter matrix size: ")
	fmt.Scan(&n)

	matrix := matrix_create(n)
	matrix_display(matrix)
	test(matrix)
}
