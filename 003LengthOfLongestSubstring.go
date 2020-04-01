package main

/**
2020-4-2
使用Map记录下字符和索引值的对于关系，在遍历的时候，如果找到了，看下这个索引值是如果是小于start（当前最大长度不重复字串的起始值）
如果大于，需要重新将当前最大长度不重复字串的起始值置为之前记录的值的下一个索引
*/
func lengthOfLongestSubstring(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	for i, ch := range []byte(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}
