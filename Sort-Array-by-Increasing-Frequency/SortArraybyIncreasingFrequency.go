// FILE: Sort-Array-by-Increasing-Frequency/SortArraybyIncreasingFrequency.go
// AUTH: ydhmkpsryk321@gmail.com
// DES: 2024-04-17 15:00:04
package leetcode

import "sort"

func frequencySort(nums []int) []int {
	freq := map[int]int{}
	for _, v := range nums {
		freq[v]++
	}
	sort.Slice(nums, func(i, j int) bool {
		if freq[nums[i]] == freq[nums[j]] {
			return nums[j] < nums[i]
		}
		return freq[nums[i]] < freq[nums[j]]
	})
	return nums
}
