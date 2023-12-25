// FILE: Create-Target-Array-in-the-Given-Order/CreateTargetArrayintheGivenOrder.go
// AUTH: ydhmkpsryk321@gmail.com
// DES: 2023-12-25 20:46:05
package leetcode

func createTargetArray(nums []int, index []int) []int {
	result := make([]int, len(nums))
	for i, pos := range index {
		copy(result[pos+1:i+1], result[pos:i])
		result[pos] = nums[i]
	}
	return result
}
