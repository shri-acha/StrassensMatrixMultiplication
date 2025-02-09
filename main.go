// package main
//
// import (
// 	"fmt"
// 	"time"
//
// )
//
// func main() {
//   bruteForceTime := make(chan time.Duration) 
//   strassenTime := make(chan time.Duration) 
//   matrix := matrix_create(512) 
//
//   defer close(bruteForceTime)
//   defer close(strassenTime)
//
//   go BenchmarkOperationBruteForceMul(bruteForceTime,matrix)
//   go BenchmarkOperationStrassenMul(strassenTime,matrix)
//
//   fmt.Printf("%v:bruteForceTime\n",<-bruteForceTime)
//   fmt.Printf("%v:strassenTime\n",<-strassenTime)
// }

package main

import "fmt"

func main() {

	var n int
	fmt.Println("Enter matrix size: ")
	fmt.Scan(&n)

	matrix1 := matrix_create(n)
	matrix2 := matrix_create(n)

  //check for power of 2
  if n>0 && n&(n-1) != 0 {
    fmatrix1 := strassen_not_x2(matrix1)
    fmatrix2 := strassen_not_x2(matrix2)
    matrix_display(fmatrix1)
    matrix_display(fmatrix2)
    
    fpro := strassen(fmatrix1,fmatrix2)
    final_product := reduced_matrix_size(fpro,n)
    matrix_display(final_product)
    return
  }

  fmt.Println("matrix1")
  matrix_display(matrix1)
  fmt.Println("matrix2")
  matrix_display(matrix2)
	pro := strassen(matrix1, matrix2)
  matrix_display(pro)

}
