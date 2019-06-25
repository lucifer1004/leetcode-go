package numenclaves

type Cell struct {
	r, c int
}

func numEnclaves(A [][]int) int {
	var directions = [4][2]int{[2]int{-1, 0}, [2]int{0, 1}, [2]int{1, 0}, [2]int{0, -1}}

	m := len(A)
	n := len(A[0])

	bfs := []Cell{}
	inQueue := map[int]bool{}

	countOnes := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if A[i][j] == 1 {
				countOnes++
			}
		}
	}

	for i := 0; i < m; i++ {
		if A[i][0] == 1 {
			bfs = append(bfs, Cell{
				r: i,
				c: 0,
			})
			inQueue[n*i] = true
		}
		if A[i][n-1] == 1 && !inQueue[n*i+n-1] {
			bfs = append(bfs, Cell{
				r: i,
				c: n - 1,
			})
			inQueue[n*i+n-1] = true
		}
	}

	for j := 0; j < n; j++ {
		if A[0][j] == 1 && !inQueue[j] {
			bfs = append(bfs, Cell{
				r: 0,
				c: j,
			})
			inQueue[j] = true
		}
		if A[m-1][j] == 1 && !inQueue[(m-1)*n+j] {
			bfs = append(bfs, Cell{
				r: m - 1,
				c: j,
			})
			inQueue[(m-1)*n+j] = true
		}
	}

	head := 0
	for head < len(bfs) {
		for _, dire := range directions {
			nextR := bfs[head].r + dire[0]
			nextC := bfs[head].c + dire[1]
			if nextR < 0 || nextR >= m || nextC < 0 || nextC >= n {
				continue
			}
			if A[nextR][nextC] == 1 && !inQueue[nextR*n+nextC] {
				bfs = append(bfs, Cell{
					r: nextR,
					c: nextC,
				})
				inQueue[nextR*n+nextC] = true
			}
		}
		head++
	}

	return countOnes - len(bfs)
}
