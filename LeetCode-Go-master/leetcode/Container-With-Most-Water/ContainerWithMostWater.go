// FILE: LeetCode-Go-master/leetcode/Container-With-Most-Water/ContainerWithMostWater.go
// AUTH: ydhmkpsryk321@gmail.com
// DES: 2024-04-28 15:57:22
package leetcode

func maxArea(height []int) int {
	max, start, end := 0, 0, len(height)-1
	for start < end {
		width := end - start
		high := 0
		if height[start] < height[end] {
			high = height[start]
			start++
		} else {
			high = height[end]
			end--
		}

		temp := width * high
		if temp > max {
			max = temp
		}
	}
	return max
}
