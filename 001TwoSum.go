package main

// 2020-3-31
// 思路：
// 用一个map存储，以数组的值为Key，以数组的索引为Value，题目中提到只有一组解，因此只要用target-num的值看下map里是否存在即可
// 这题是简单难度，主要还是能习惯下Golang里map的基本用法
func twoSum(nums []int, target int) []int {
	ret := make([]int, 2)
	numMaps := map[int]int{}
	for idx, num := range nums {
		if v, ok := numMaps[target-num]; ok {
			ret[0] = v
			ret[1] = idx
		} else {
			numMaps[num] = idx
		}
	}
	return ret
}
