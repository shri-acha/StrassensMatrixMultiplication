package main

func matrix_bruteforce(m1 [][]int, m2 [][]int) [][]int {
	var i, j, k int
	n := len(m1)
	m3 := make([][]int, n) // Create n rows

	for i = range m3 {
		m3[i] = make([]int, n) // Iterate over n rows and create n columns per row
	}
	sum := 0

	for i = 0; i < n; i++ {
		for j = 0; j < n; j++ {
			for k = 0; k < n; k++ {
				pro := m1[i][k] * m2[k][j]
				sum = sum + pro
			}
			m3[i][j] = sum
			sum = 0
		}
	}
	return m3
}
