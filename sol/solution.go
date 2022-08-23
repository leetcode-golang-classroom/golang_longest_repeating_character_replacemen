package sol

func characterReplacement(s string, k int) int {
	freq, sLen := make(map[byte]int), len(s)
	left, maxLen, count := 0, 0, 0
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for right := 0; right < sLen; right++ {
		freq[s[right]]++
		if freq[s[right]] > count {
			count = freq[s[right]]
		}
		// slide window over size
		if right-left+1-count > k {
			freq[s[left]]--
			left++
		}
		maxLen = max(maxLen, right-left+1)
	}
	return maxLen
}
