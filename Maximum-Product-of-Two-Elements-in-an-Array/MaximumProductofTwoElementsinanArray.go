// FILE: Maximum-Product-of-Two-Elements-in-an-Array/MaximumProductofTwoElementsinanArray.go
// AUTH: ydhmkpsryk321@gmail.com
// DES: 2024-03-23 10:29:27
package leetcode

func maxProduct(nums []int) int {
	max1, max2 := 0, 0
	for _, num := range nums {
		if num >= max1 {
			max2 = max1
			max1 = num
		} else if num <= max1 && num >= max2 {
			max2 = num
		}
	}
	return (max1 - 1) * (max2 - 1)
}
