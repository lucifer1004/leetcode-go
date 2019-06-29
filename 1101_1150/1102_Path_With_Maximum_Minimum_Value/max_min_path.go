package maxminpath

import (
	"container/heap"
)

type IntHeap []int

var mat []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return mat[h[i]] > mat[h[j]] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maximumMinimumPath(A [][]int) int {
	r := len(A)
	c := len(A[0])
	d := [4][2]int{[2]int{-1, 0}, [2]int{0, 1}, [2]int{1, 0}, [2]int{0, -1}}

	if r == 1 && c == 1 {
		return A[0][0]
	}

	mat = []int{}

	n := r * c
	for i := 0; i < n; i++ {
		mat = append(mat, 0)
	}

	for k := range d {
		m1 := d[k][0]
		n1 := d[k][1]
		if m1 < 0 || m1 >= r || n1 < 0 || n1 >= c {
			continue
		}
		next := m1*c + n1
		val := min(A[0][0], A[m1][n1])
		mat[next] = val
	}

	list := &IntHeap{}
	heap.Init(list)
	inList := map[int]bool{}
	for i := 1; i < n; i++ {
		if mat[i] > 0 {
			heap.Push(list, i)
			inList[i] = true
		}
	}

	for list.Len() > 0 {
		head := heap.Pop(list).(int)
		m := head / c
		n := head % c
		for k := range d {
			m1 := m + d[k][0]
			n1 := n + d[k][1]
			if m1 < 0 || m1 >= r || n1 < 0 || n1 >= c {
				continue
			}
			i := m1*c + n1
			if i == 0 {
				continue
			}
			newVal := min(mat[head], min(A[m1][n1], A[m][n]))
			if newVal > mat[i] {
				mat[i] = newVal
				if !inList[i] {
					heap.Push(list, i)
					inList[i] = true
				}
			}
		}
	}

	return mat[n-1]
}
