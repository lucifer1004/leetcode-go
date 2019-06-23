package brackexpansion

import "sort"

func deduplicate(arr []string) []string {
	flag := make([]bool, len(arr))
	for i := range flag {
		flag[i] = true
	}

	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] == arr[i] {
				flag[j] = false
			}
		}
	}

	ans := []string{}
	for i, v := range flag {
		if v {
			ans = append(ans, arr[i])
		}
	}

	return ans
}

func braceExpansion(s string) []string {
	ops := []string{}
	strs := [][]string{}
	s = "{" + s + "}"
	curr := ""
	for i := range s {
		char := string(s[i])
		if char == "{" {
			if curr != "" {
				strs = append(strs, []string{curr})
				curr = ""
			}

			// Add "*" if necessary
			if i > 0 {
				prevChar := string(s[i-1])
				if prevChar != "," && prevChar != "{" {
					ops = append(ops, "*")
				}
			}
			ops = append(ops, char)
			continue
		}

		if char == "," {
			if curr != "" {
				strs = append(strs, []string{curr})
				curr = ""
			}

			// Deal with remaining "*"s
			j := len(ops) - 1
			for ops[j] == "*" {
				if len(strs) > 1 {
					k := len(strs) - 1
					tmp := []string{}
					for _, a := range strs[k-1] {
						for _, b := range strs[k] {
							tmp = append(tmp, a+b)
						}
					}
					strs[k-1] = tmp
					strs = strs[0:k]
				}
				j--
			}
			ops = ops[0 : j+1]
			ops = append(ops, char)
			continue
		}

		// Deal with ","s and "*"s within a pair of "{}"
		if char == "}" {
			if curr != "" {
				strs = append(strs, []string{curr})
				curr = ""
			}

			j := len(ops) - 1
			for ops[j] != "{" {
				if ops[j] == "," {
					k := len(strs) - 1
					if k > 0 {
						strs[k-1] = append(strs[k-1], strs[k]...)
						strs = strs[0:k]
					}
				}

				if ops[j] == "*" {
					if len(strs) > 1 {
						k := len(strs) - 1
						tmp := []string{}
						for _, a := range strs[k-1] {
							for _, b := range strs[k] {
								tmp = append(tmp, a+b)
							}
						}
						strs[k-1] = tmp
						strs = strs[0:k]
					}
				}
				j--
			}
			ops = ops[0:j]

			// Add "*" if necessary
			if i < len(s)-1 {
				nextChar := string(s[i+1])
				if nextChar != "}" && nextChar != "{" && nextChar != "," {
					ops = append(ops, "*")
				}
			}
			continue
		}
		curr += char
	}

	return strs[0]
}

func braceExpansionII(expression string) []string {
	ans := deduplicate(braceExpansion(expression))
	sort.Strings(ans)
	return ans
}
