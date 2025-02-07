package main

func matrix_split(matrix [][]int) ([][]int, [][]int, [][]int, [][]int) {

	n := len(matrix) // len(matrix) returns the number of rows
	n_half := n / 2  // Splitting matrix into half its size

	matrix_1 := make([][]int, n_half)
	matrix_2 := make([][]int, n_half)
	matrix_3 := make([][]int, n_half)
	matrix_4 := make([][]int, n_half)

	for i := 0; i < n_half; i++ {
		matrix_1[i] = matrix[i][:n_half] // Top left :- For example, for 4x4 matrix[i] represents matrix[0] and matrix[1] here, indicating the first two rows.
		//[:n_half] iterates from first column to the last column i.e 4/2 = 2nd column to give the corresponding 2 columns. Same thing done in the
		//subsequent examples
		matrix_2[i] = matrix[i][n_half:]        // Top right
		matrix_3[i] = matrix[n_half+i][:n_half] //Bottom left
		matrix_4[i] = matrix[n_half+i][n_half:] //Bottom right
	}

	return matrix_1, matrix_2, matrix_3, matrix_4

}
