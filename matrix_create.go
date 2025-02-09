package main


func matrix_create(n int) (matrix [][]int) {
	matrix = make([][]int, n) // Create n rows, columns not allocated
  for i := range matrix {
		matrix[i] = make([]int, n) //Iterate over n rows and create n columns per row
	}
	for i := 0; i < n; i++ {
    counter := 0
		for j := 0; j < n; j++ {

			matrix[i][j] = counter
			counter++

		}
	}

	return matrix
}
