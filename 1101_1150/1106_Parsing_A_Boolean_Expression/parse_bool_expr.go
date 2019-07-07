package parseboolexpr

func getSub(expression string) []string {
	depth := 0
	start := 2
	sub := []string{}
	for i := 1; i < len(expression); i++ {
		char := string(expression[i])
		if char == "(" {
			depth++
		}
		if char == ")" {
			depth--
			if depth == 0 {
				sub = append(sub, expression[start:i])
			}
		}
		if char == "," && depth == 1 {
			sub = append(sub, expression[start:i])
			start = i + 1
		}
	}
	return sub
}

func parseBoolExpr(expression string) bool {
	if expression == "t" {
		return true
	}
	if expression == "f" {
		return false
	}
	symbol := expression[0:1]
	if symbol == "!" {
		return !parseBoolExpr(expression[2 : len(expression)-1])
	}
	if symbol == "&" {
		sub := getSub(expression)
		ans := parseBoolExpr(sub[0])
		for i := 1; i < len(sub); i++ {
			ans = ans && parseBoolExpr(sub[i])
		}
		return ans
	}
	if symbol == "|" {
		sub := getSub(expression)
		ans := parseBoolExpr(sub[0])
		for i := 1; i < len(sub); i++ {
			ans = ans || parseBoolExpr(sub[i])
		}
		return ans
	}
	return false
}
