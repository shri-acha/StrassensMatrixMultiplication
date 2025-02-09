package main

import "time"

func BenchmarkOperationStrassenMul(strassenTime chan time.Duration,matrix [][]int) {
  start := time.Now()
  strassen(matrix,matrix)
  strassenTime<-time.Since(start)
}
func BenchmarkOperationBruteForceMul(bruteForceTime chan time.Duration,matrix [][]int) {
  start := time.Now()
  matrix_bruteforce(matrix,matrix)
  bruteForceTime<-time.Since(start)
}

