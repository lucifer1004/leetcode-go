package isinterleave

func isInterleave(s1 string, s2 string, s3 string) bool {
	if (len(s1) == 0 && s2 == s3) || (len(s2) == 0 && s1 == s3) {
		return true
	}

	if len(s1)+len(s2) != len(s3) {
		return false
	}

	ans := false
	if len(s1) > 0 && s1[0] == s3[0] {
		ans = isInterleave(s1[1:], s2, s3[1:])
	}

	if len(s2) > 0 && s2[0] == s3[0] {
		ans = ans || isInterleave(s1, s2[1:], s3[1:])
	}

	return ans
}
