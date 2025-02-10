// package main
// // import "fmt"

// func strassen_not_x2(m [][]int) [][]int {

//   n:= len(m)

//   n= nearest_power_of_2(n)

// 	final_matrix := make([][]int, n)

//   for i := range final_matrix {
// 		final_matrix[i] = make([]int, n) //To understand how matrix
// 		//is made here, refer to matrix_create.go
// 	}

//   for i:=0;i<len(m);i++ {
//     for j:=0;j<len(m);j++ {
//       final_matrix[i][j] = m[i][j]
//     }
//   }

//   return final_matrix
// }

package main

func padding(m [][]int, l int) [][]int {

	final_matrix := make([][]int, l)

	for i := range final_matrix {
		final_matrix[i] = make([]int, l)
	}

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			final_matrix[i][j] = m[i][j]
		}
	}

	return final_matrix
}
