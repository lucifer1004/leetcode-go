package ismatch

func isMatch(s string, p string) bool {
	if s == p {
		return true
	}

	l1 := len(s)
	l2 := len(p)
	if (l1 > 0 && l2 == 0) || string(p[0]) == "*" {
		return false
	}

	mat := [][]bool{}
	for i := 0; i <= 1; i++ {
		arr := make([]bool, l2+1)
		mat = append(mat, arr)
	}

	mat[1][0] = true

	for i := 0; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if i == 0 {
				if string(p[j-1]) == "*" {
					last := string(p[j-2])
					if last == "*" {
						mat[1][j] = mat[1][j-1]
					} else {
						mat[1][j] = mat[1][j-2]
					}
				}
				continue
			}

			ans := false
			if string(p[j-1]) == "." || p[j-1] == s[i-1] {
				ans = mat[0][j-1]
			}
			if string(p[j-1]) == "*" {
				ans = mat[1][j-1] || mat[1][j-2]

				last := string(p[j-2])
				if last == "." || last == string(s[i-1]) {
					ans = ans || mat[0][j] || mat[0][j-1]
				}
			}
			mat[1][j] = ans
		}
		mat[0] = mat[1]
		mat[1] = make([]bool, l2+1)
	}
	return mat[0][l2]
}
