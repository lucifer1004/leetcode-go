package isscramble

func eq(s1, s2 string) bool {
	count := make([]int, 300)

	for _, v := range s1 {
		count[int(v)]++
	}

	for _, v := range s2 {
		count[int(v)]--
	}

	for _, v := range count {
		if v != 0 {
			return false
		}
	}

	return true
}

func isScramble(s1 string, s2 string) bool {
	var ans bool

	if s1 == s2 {
		return true
	}

	if !eq(s1, s2) {
		return false
	}

	l := len(s1)
	for i := 1; i < l; i++ {
		l1 := s1[0:i]
		r1 := s1[i:l]
		l21 := s2[0:i]
		r21 := s2[i:l]
		l22 := s2[0 : l-i]
		r22 := s2[l-i : l]
		ans = ans || (isScramble(l1, l21) && isScramble(r1, r21)) || (isScramble(l1, r22) && isScramble(r1, l22))
	}

	return ans
}
