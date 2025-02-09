package main

func reduced_matrix_size(m [][]int, n int) [][]int {

  final_matrix := make([][]int, n)
  
  for i := range final_matrix {
		final_matrix[i] = make([]int, n) //To understand how matrix
		//is made here, refer to matrix_create.go
	}
  for i:=0;i<n;i++ {
    for j:=0;j<n;j++ {
      final_matrix[i][j] = m[i][j]
    }
  }
  return final_matrix
}
