package minwindow

// Store positive value of count to ref
func makeRef(a []int) []int {
	ref := []int{}
	for i, v := range a {
		if v > 0 {
			ref = append(ref, i)
		}
	}
	return ref
}

// Compare two counts according to ref
func gt(a, b, ref []int) bool {
	for _, v := range ref {
		if a[v] < b[v] {
			return false
		}
	}

	return true
}

// Count character frequency
func strToCount(s string) []int {
	count := make([]int, 300)
	for _, v := range s {
		count[byte(v)]++
	}
	return count
}

// Calculate the minimum window substring
func minWindow(s string, t string) string {
	// Target is empty
	if t == "" {
		return ""
	}

	tc := strToCount(t)
	sc := strToCount(s)
	ref := makeRef(tc)

	// Task is impossible
	if !gt(sc, tc, ref) {
		return ""
	}

	// Filter characters in source according to target
	p := []int{}
	for i := range s {
		if tc[s[i]] > 0 {
			p = append(p, i)
		}
	}

	// Initialize pointers
	c := make([]int, 300)
	l := 0
	r := 0
	c[s[p[r]]]++
	ans := s

	for r < len(p) {
		if gt(c, tc, ref) {
			// Found a valid solution
			// Update current answer if the new solution is better
			if p[r]-p[l]+1 < len(ans) {
				ans = s[p[l] : p[r]+1]
			}
			// Move left pointer
			c[s[p[l]]]--
			l++
		} else {
			// Move right pointer
			r++
			if r < len(p) {
				c[s[p[r]]]++
			}
		}
	}

	return ans
}
