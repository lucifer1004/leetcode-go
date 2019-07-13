package removevowels

import "regexp"

func removeVowels(s string) string {
	var re = regexp.MustCompile(`(a|e|i|o|u)`)
	return re.ReplaceAllString(s, ``)
}
