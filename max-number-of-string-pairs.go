package leetcodesolutions

func maximumNumberOfStringPairs(words []string) int {
	count := 0
	for i := 0; i < len(words)-1; i++ {
		for j := i + 1; j < len(words); j++ {
			if isReversed(words[i], words[j]) {
				count++
			}
		}
	}
	return count
}

func isReversed(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	right := 0
	left := len(str1) - 1
	for i := 0; i < len(str1); i++ {
		if str1[right] != str2[left] {
			return false
		}
		right++
		left--
	}
	return true
}
