package leetcodesolutions

func isMatch(s string, p string) bool {

	if len(p) == 0 {
		return len(s) == 0
	}

	firstCharMatch := len(s) > 0 && (p[0] == s[0] || p[0] == '.')

	if len(p) >= 2 && p[1] == '*' {
		return isMatch(s, p[2:]) || (firstCharMatch && isMatch(s[1:], p))
	} else {
		return firstCharMatch && isMatch(s[1:], p[1:])
	}

}
