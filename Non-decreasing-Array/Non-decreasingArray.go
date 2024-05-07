// FILE: Non-decreasing-Array/Non-decreasingArray.go
// AUTH: ydhmkpsryk321@gmail.com
// DES: 2024-05-07 17:12:26
package leetcode

func checkPossibility(nums []int) bool {
	count := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			count++
			if count > 1 {
				return false
			}
			if i > 0 && nums[i+1] < nums[i-1] {
				nums[i+1] = nums[i]
			}
		}
	}
	return true
}
