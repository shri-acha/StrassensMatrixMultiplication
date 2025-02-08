package main

import "fmt"

func main() {

	var n int
	fmt.Println("Enter matrix size: ")
	fmt.Scan(&n)

	matrix1 := matrix_create(n)
	matrix2 := matrix_create(n)
	matrix_display(matrix1)
	matrix_display(matrix2)

	pro := strassen(matrix1, matrix2)
	matrix_display(pro)

}
