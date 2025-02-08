//required for strassen
//cant subtract slices directly unlike python

package main

func matrix_sub(matrices ...[][]int) [][]int {
	n := len(matrices[0])
	diff := make([][]int, n)

	for i := range diff {
		diff[i] = make([]int, n)
		copy(diff[i], matrices[0][i]) // Start with first matrix
	}

	for _, matrix := range matrices[1:] { // Start subtraction from second matrix
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				diff[i][j] -= matrix[i][j]
			}
		}
	}

	return diff
}
