package numtilepossibilities

func lt(a, b []int) bool {
	for i := range a {
		if a[i] < b[i] {
			return true
		}
	}
	return false
}

func addOne(a, upper []int) {
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] < upper[i] {
			a[i]++
			for j := i + 1; j < len(a); j++ {
				a[j] = 0
			}
			break
		}
	}
}

func calc(a []int) int {
	fact := [8]int{1, 1, 2, 6, 24, 120, 720, 5040}
	sum := 0
	for _, v := range a {
		sum += v
	}
	ans := fact[sum]
	for _, v := range a {
		if v > 1 {
			ans /= fact[v]
		}
	}
	return ans
}

func numTilePossibilities(tiles string) int {
	freq := map[byte]int{}
	upper := []int{}
	for _, v := range tiles {
		freq[byte(v)]++
	}
	for _, v := range freq {
		upper = append(upper, v)
	}

	ans := 0
	curr := make([]int, len(upper))
	for lt(curr, upper) {
		addOne(curr, upper)
		ans += calc(curr)
	}

	return ans
}
