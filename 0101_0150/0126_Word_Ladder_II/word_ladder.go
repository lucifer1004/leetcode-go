package wordladder

func isValid(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	count := 0
	for i := range a {
		if a[i] != b[i] {
			count++
		}
	}

	if count == 1 {
		return true
	}

	return false
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	found := false
	end := -1
	for i, v := range wordList {
		if v == endWord {
			found = true
			end = i
			break
		}
	}
	if !found {
		return [][]string{}
	}

	found = false
	idx := 0
	for i, v := range wordList {
		if v == beginWord {
			found = true
			idx = i
			break
		}
	}
	if !found {
		wordList = append(wordList, beginWord)
		idx = len(wordList) - 1
	}

	mat := [][]bool{}
	for i := 0; i < len(wordList); i++ {
		arr := make([]bool, len(wordList))
		mat = append(mat, arr)
	}

	for i := 0; i < len(wordList)-1; i++ {
		for j := i + 1; j < len(wordList); j++ {
			if isValid(wordList[i], wordList[j]) {
				mat[i][j] = true
				mat[j][i] = true
			}
		}
	}

	list := []int{}
	dist := make([][][]int, len(wordList))
	for i := range wordList {
		if mat[idx][i] {
			list = append(list, i)
			dist[i] = [][]int{[]int{i}}
		} else {
			dist[i] = [][]int{}
		}
	}

	head := 0
	for head < len(list) {
		from := list[head]
		for to := range wordList {
			if to == idx {
				continue
			}
			if from != to && mat[from][to] {
				if len(dist[to]) == 0 {
					list = append(list, to)
				}

				if len(dist[to]) == 0 || len(dist[from][0])+1 <= len(dist[to][0]) {
					if len(dist[to]) == 0 || len(dist[from][0])+1 < len(dist[to][0]) {
						dist[to] = [][]int{}
					}

					for _, v := range dist[from] {
						path := make([]int, len(v))
						copy(path, v)
						path = append(path, to)
						dist[to] = append(dist[to], path)
					}
				}
			}
		}
		head++
	}

	ans := [][]string{}

	for _, v := range dist[end] {
		res := []string{wordList[idx]}
		for _, i := range v {
			res = append(res, wordList[i])
		}
		ans = append(ans, res)
	}

	return ans
}
