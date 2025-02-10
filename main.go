package main

import (
	"fmt"
	"time"
)

func main() {
  bruteForceTime := make(chan time.Duration) 
  strassenTime := make(chan time.Duration) 
  matrix := matrix_create(3) 

  defer close(bruteForceTime)
  defer close(strassenTime)

  go BenchmarkOperationBruteForceMul(bruteForceTime,matrix)
  go BenchmarkOperationStrassenMul(strassenTime,matrix)

  fmt.Printf("%v:bruteForceTime\n",<-bruteForceTime)
  fmt.Printf("%v:strassenTime\n",<-strassenTime)
}
