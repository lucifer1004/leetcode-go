package main

import "fmt"

func isInterleave(s1 string, s2 string, s3 string) bool {
	l1 := len(s1)
	l2 := len(s2)
	l3 := len(s3)

	if l1+l2 != l3 {
		return false
	}

	mat := [][]bool{}
	for i := 0; i <= l1; i++ {
		arr := make([]bool, l2+1)
		mat = append(mat, arr)
	}

	for i := 0; i <= l1; i++ {
		for j := 0; j <= l2; j++ {
			if i == 0 {
				if s2[0:j] == s3[0:j] {
					mat[i][j] = true
				}
				continue
			}

			if j == 0 {
				if s1[0:i] == s3[0:i] {
					mat[i][j] = true
				}
				continue
			}

			mat[i][j] = (s1[i-1] == s3[i+j-1] && mat[i-1][j]) || (s2[j-1] == s3[i+j-1] && mat[i][j-1])
		}
	}

	return mat[l1][l2]
}

func main() {
	fmt.Println(isInterleave(
		"aa",
		"bb",
		"abab",
	))
}
