package restoreipaddresses

import "strconv"

func isValid(s string) bool {
	if string(s[0]) == "0" && len(s) > 1 {
		return false
	}

	d, _ := strconv.Atoi(s)
	if d >= 0 && d <= 255 {
		return true
	}

	return false
}

func restore(s string, n int) []string {
	ans := []string{}
	if n == 1 {
		if isValid(s) {
			return []string{s}
		}
		return ans
	}

	l := len(s)
	u := 3

	if l < n {
		return ans
	}
	if l-n+1 < u {
		u = l - n + 1
	}

	for i := 1; i <= u; i++ {
		if isValid(s[0:i]) {
			subAns := restore(s[i:], n-1)
			for _, v := range subAns {
				ans = append(ans, s[0:i]+"."+v)
			}
		}
	}

	return ans
}

func restoreIPAddresses(s string) []string {
	return restore(s, 4)
}
