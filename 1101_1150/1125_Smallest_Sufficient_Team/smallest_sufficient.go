package smallestsufficient

func sum(a []int) int {
	ans := 0
	for _, v := range a {
		ans += v
	}
	return ans
}

func check(curr, target int, used, p, r []int) bool {
	ans := curr
	for i := range used {
		if used[i] == 1 {
			ans |= p[r[i]]
		}
	}
	return ans == target
}

func solve(curr, target int, p, r []int) []int {
	min := len(r) + 1
	ans := []int{}
	upper := 1
	for range r {
		upper <<= 1
	}
	for i := 0; i < upper; i++ {
		used := make([]int, len(r))
		t := i
		idx := 0
		for t > 0 {
			used[idx] = t % 2
			t >>= 1
			idx++
		}
		if check(curr, target, used, p, r) && sum(used) < min {
			min = sum(used)
			newAns := []int{}
			for j := range used {
				if used[j] == 1 {
					newAns = append(newAns, r[j])
				}
			}
			ans = newAns
		}
	}
	return ans
}

func smallestSufficientTeam(reqSkills []string, people [][]string) []int {
	skillMap := map[string]int{}
	skill := 1
	for _, v := range reqSkills {
		skillMap[v] = skill
		skill <<= 1
	}
	skill--
	p := make([]int, len(people))
	for i := range people {
		for j := range people[i] {
			p[i] += skillMap[people[i][j]]
		}
	}
	ans := []int{}
	curr := 0

	// Pick people with unique skills
	for i := range p {
		rest := 0
		for j := range p {
			if j == i {
				continue
			}
			rest |= p[j]
		}
		if rest != skill {
			ans = append(ans, i)
		}
	}

	for _, v := range ans {
		curr |= p[v]
	}

	if curr == skill {
		return ans
	}

	// Exclude people overlapped by others
	flag := make([]bool, len(p))
	for _, v := range ans {
		flag[v] = true
	}
	for i := range p {
		for j := i + 1; j < len(p); j++ {
			common := p[i] & p[j]
			if p[i] == p[j] {
				flag[j] = true
				continue
			}
			if p[i] == common {
				flag[i] = true
				continue
			}
			if p[j] == common {
				flag[j] = true
			}
		}
	}
	r := []int{}
	for i := range flag {
		if !flag[i] {
			r = append(r, i)
		}
	}

	ans = append(ans, solve(curr, skill, p, r)...)

	return ans
}
