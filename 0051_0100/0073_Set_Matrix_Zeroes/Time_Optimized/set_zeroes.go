package setzeroes

func setZeroes(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])
	rows := map[int]bool{}
	cols := map[int]bool{}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				rows[i] = true
				cols[j] = true
			}
		}
	}

	for k := range rows {
		for j := 0; j < n; j++ {
			matrix[k][j] = 0
		}
	}

	for k := range cols {
		for i := 0; i < m; i++ {
			matrix[i][k] = 0
		}
	}
}
