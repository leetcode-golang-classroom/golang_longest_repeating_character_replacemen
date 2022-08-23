# golang_longest_repeating_character_replacement

You are given a string s and an integer k. You can choose any character of the string and change it to any other uppercase English character. You can perform this operation at most k times.

Return the length of the longest substring containing the same letter you can get after performing the above operations.

## Examples

**Example 1:**

```
Input: s = "ABAB", k = 2
Output: 4
Explanation: Replace the two 'A's with two 'B's or vice versa.

```

**Example 2:**

```
Input: s = "AABABBA", k = 1
Output: 4
Explanation: Replace the one 'A' in the middle with 'B' and form "AABBBBA".
The substring "BBBB" has the longest repeating letters, which is 4.

```

**Constraints:**

- `1 <= s.length <= 105`
- `s` consists of only uppercase English letters.
- `0 <= k <= s.length`

## 解析

給定一個字串 s, 還有一個非負整數 k 

其中每個 s[i] 都是大寫英文字母

要求寫一個演算法字串 s 中在可以允許更換 k 次字元的情況下，找出最長的字元不重複字元的子字串長度

要找出不重複字元的最長子字串，相當於是要是找出出現最多次的字元並且每個字元位置必須相連

找出出現最多的字元個數可以透過 hashTable 的方式把每個遇到的字元個數累加起來

而位置相連的部份則可以透過 slide window 的方式來處理

定義一個左界為 left = 0, 右界right 從 0 開始逐步往右移 

每次遇到 更新當下  freq[s[right]]= freq[s[right]]+1, 並且把最大的 freq 紀錄在一個 count 次數內

![](https://i.imgur.com/giebUzr.png)


因為可以更換字元 k 次, 代表只要在slide window 長度(右界 -左界 + 1) - 目前最大累計字元個數 ≤ k 都可以 替換為同一個字元

當 slide window 長度 - 最大累計字元個數 > k 時， 代表無法在把原本字元延長。

所以需要把左界的字元次數 - 1, 左界向右移

![](https://i.imgur.com/XUzA9WN.png)


每次更新 maxLen = max(maxLen, left - right +1)

最後回傳 maxLen

具體作法如下

![](https://i.imgur.com/xUKqjPW.png)

## 程式碼
```go
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

```
## 困難點

1. 需要看出最長相同字元的子字串形成條件
2. 需要知道 slide window 的移動條件，何時需要把左界向右

## Solve Point

- [x]  初始化 left = 0, maxLen = 0, count = 0, freq 為一個 hashMap 用來紀錄每個字元出現的次數
- [x]  當 right = 0.. len(s)-1 時做以下運算
- [x]  更新 freq[s[right]]= freq[s[right]]+1
- [x]  當 freq[s[right]] > count 時，更新 count = freq[s[right]]
- [x]  當 right - left + 1 - count > k 時， 更新 freq[s[left]] = freq[s[left]] - 1, left++
- [x]  更新 maxLen  = max(maxLen, right - left + 1)
- [x]  回傳 maxLen