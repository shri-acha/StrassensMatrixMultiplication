package main

import "fmt"

func matrix_display(matrix [][]int) {
	// println("Generated Matrix: ")
	for _, row := range matrix {
		fmt.Println(row)
	}
	fmt.Printf("\n")
}
