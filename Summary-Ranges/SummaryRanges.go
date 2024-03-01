// FILE: Summary-Ranges/SummaryRanges.go
// AUTH: ydhmkpsryk321@gmail.com
// DES: 2024-03-01 15:20:11
package leetcode

import (
	"strconv"
)

func summaryRanges(nums []int) (ans []string) {
	for i, n := 0, len(nums); i < n; {
		left := i
		for i++; i < n && nums[i-1]+1 == nums[i]; i++ {
		}
		s := strconv.Itoa(nums[left])
		if left != i-1 {
			s += "->" + strconv.Itoa(nums[i-1])
		}
		ans = append(ans, s)
	}
	return
}
