package main

func matrix_create(n int) (matrix [][]int) {
	matrix = make([][]int, n) // Create n rows, columns not allocated
	counter := 1
	for i := range matrix {
		matrix[i] = make([]int, n) //Iterate over n rows and create n columns per row
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j] = 1
			counter++
		}
	}

	return matrix
}
