package main

// import "fmt"

func strassen(m1 [][]int, m2 [][]int) [][]int {
	n := len(m1)
  const LEAF_SIZE = 8
	final_matrix := make([][]int, n)

	for i := range final_matrix {
		final_matrix[i] = make([]int, n) //To understand how matrix
		//is made here, refer to matrix_create.go
	}

	if n <= LEAF_SIZE {
		return matrix_bruteforce(m1, m2)
	}

	a, b, c, d := matrix_split(m1)
	e, f, g, h := matrix_split(m2)

	ad := matrix_add(a, d)
	eh := matrix_add(e, h)
	p1 := strassen(ad, eh) // p1 = (a + d) * (e + h)

	ge := matrix_sub(g, e)
	p2 := strassen(d, ge) // p2 = d * (g - e)

	ab := matrix_add(a, b)
	p3 := strassen(ab, h) // p3 = (a + b) * h

	bd := matrix_sub(b, d)
	gh := matrix_add(g, h)
	p4 := strassen(bd, gh) // p4 = (b - d) * (g + h)

	fh := matrix_sub(f, h)
	p5 := strassen(a, fh) // p5 = a * (f - h)

	cd := matrix_add(c, d)
	p6 := strassen(cd, e) // p6 = (c + d) * e

	ac := matrix_sub(a, c)
	ef := matrix_add(e, f)
	p7 := strassen(ac, ef) // p7 = (a - c) * (e + f)

	// fmt.Println("P0:", p1)
	// fmt.Println("P2:", p2)
	// fmt.Println("P3:", p3)
	// fmt.Println("P4:", p4)
	// fmt.Println("P5:", p5)
	// fmt.Println("P6:", p6)
	// fmt.Println("P7:", p7)

	c00 := matrix_sub(matrix_add(p1, p2, p4), p3)
	// c00 = p1 + p2 - p3 + p4
	//Note that matrix add and matrix sub can take any number of matrices
	//to be added or subtracted

	c01 := matrix_add(p5, p3)
	// c01 = p3 + p5

	c10 := matrix_add(p2, p6)
	// c10 = p2 + p6

	c11 := matrix_sub(matrix_add(p1, p5), matrix_add(p6, p7))
	// c11 = p1 + p5 - p6 - p7

	//Top-left
	for i := 0; i < n/2; i++ {
		for j := 0; j < n/2; j++ {
			final_matrix[i][j] = c00[i][j]
		}
	}

	//Top-right
	for i := 0; i < n/2; i++ {
		for j := 0; j < n/2; j++ {
			final_matrix[i][j+n/2] = c01[i][j]
		}
	}

	//Bottom-left
	for i := 0; i < n/2; i++ {
		for j := 0; j < n/2; j++ {
			final_matrix[i+n/2][j] = c10[i][j]
		}
	}

	//Bottom-right
	for i := 0; i < n/2; i++ {
		for j := 0; j < n/2; j++ {
			final_matrix[i+n/2][j+n/2] = c11[i][j]
		}
	}

	return final_matrix

}
