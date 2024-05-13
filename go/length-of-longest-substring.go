package leetcodesolutions

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	maxCount := 0
	charMap := make(map[byte]int)
	right, left := 0, 0

	for right < len(s) {
		_, ok := charMap[s[right]]
		if !ok {
			charMap[s[right]] = right
		} else {
			if charMap[s[right]] >= left {
				left = charMap[s[right]] + 1
			}
			charMap[s[right]] = right

		}
		right++
		dist := right - left
		if dist > maxCount {
			maxCount = dist
		}
	}
	return maxCount
}
