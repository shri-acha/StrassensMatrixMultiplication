package main

func matrix_create(r, c int) (matrix [][]int) {
	
	matrix = make([][]int, r) // Create r rows, columns not allocated
	
	for i := range matrix {
		matrix[i] = make([]int, c) //Iterate over r rows and create c columns per row
	}

	for i := 0; i < r; i++ {
		
		counter := 1

		for j := 0; j < c; j++ {

			matrix[i][j] = counter
			counter++

		}
	}

	return matrix
}
