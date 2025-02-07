package main

func matrix_split(matrix [][]int) {

	rows := len(matrix)    // len(matrix) returns the number of rows
	cols := len(matrix[0]) // Note that matrix[0] gives the first row and len(matrix[0])
	//gives number of columns in first row.

	row_half := rows / 2
	col_half := cols / 2

	matrix_1 := make([][]int, row_half)
	for i := 0; i < row_half; i++ {
		matrix_1[i] = matrix[i][:col_half]
	}

	matrix_2 := make([][]int, row_half)
	for i := 0; i < row_half; i++ {
		matrix_2[i] = matrix[i][col_half:]
	}

	matrix_3 := make([][]int, row_half)
	for i := 0; i < row_half; i++ {
		matrix_3[i] = matrix[row_half+i][:col_half]
	}

	matrix_4 := make([][]int, row_half)
	for i := 0; i < row_half; i++ {
		matrix_4[i] = matrix[row_half+i][col_half:]
	}

	matrix_display(matrix_1)
	matrix_display(matrix_2)
	matrix_display(matrix_3)
	matrix_display(matrix_4)
}
