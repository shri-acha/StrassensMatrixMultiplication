package main

import (
	"fmt"
	"time"

)

func main() {
  bruteForceTime := make(chan time.Duration)
  strassenTime := make(chan time.Duration)

  matrix := matrix_create(3,3) // example matrix say,training data and also weights corresponding to the data.
  
  go BenchmarkOperationBruteForceMul(bruteForceTime,matrix) // Using bruteforce multiplication
  go BenchmarkOperationStrassenMul(strassenTime,matrix) // Using strassen multiplication

  fmt.Printf("%v:strassenTime\n",<-strassenTime)
  fmt.Printf("%v:bruteForceTime\n",<-bruteForceTime)
}

