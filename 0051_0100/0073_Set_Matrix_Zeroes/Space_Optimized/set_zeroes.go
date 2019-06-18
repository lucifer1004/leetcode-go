package setzeros

const inf = 2147483647

func setZeroes(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				// Mark the cells to be filled with inf
				for k := 0; k < n; k++ {
					if matrix[i][k] != 0 {
						matrix[i][k] = inf
					}
				}

				for k := 0; k < m; k++ {
					if matrix[k][j] != 0 {
						matrix[k][j] = inf
					}
				}
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == inf {
				matrix[i][j] = 0
			}
		}
	}
}
