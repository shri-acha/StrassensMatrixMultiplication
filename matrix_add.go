//required for strassen
//cant add slices directly unlike python

package main

func matrix_add(matrices ...[][]int) [][]int { // The parameter indicates that it receives a variable amount of matrices.
	n := len(matrices[0]) // Calculating dimension of nxn matrix
	//Note than len(matrices) gives number of parameters passed into the function
	sum := make([][]int, n)

	for i := range sum {
		sum[i] = make([]int, n) //Iterate over n rows and create n columns per row
	}

	for _, matrix := range matrices {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				sum[i][j] += matrix[i][j]
			}
		}
	}

	return sum
}
