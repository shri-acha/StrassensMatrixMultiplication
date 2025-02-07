package main

import "fmt"

func main() {

	var n int
	fmt.Println("Enter matrix size: ")
	fmt.Scan(&n)

	matrix := matrix_create(n)
	matrix_display(matrix)
	m1, m2, m3, m4 := matrix_split(matrix)

	matrix_display(m1)
	matrix_display(m2)
	matrix_display(m3)
	matrix_display(m4)
}
