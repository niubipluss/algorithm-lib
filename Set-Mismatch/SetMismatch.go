// FILE: Set-Mismatch/SetMismatch.go
// AUTH: ydhmkpsryk321@gmail.com
// DES: 2024-04-14 17:11:39
package leetcode

func findErrorNums(nums []int) []int {
	m, res := make([]int, len(nums)), make([]int, 2)
	for _, n := range nums {
		if m[n-1] == 0 {
			m[n-1] = 1
		} else {
			res[0] = n
		}
	}
	for i := range m {
		if m[i] == 0 {
			res[1] = i + 1
			break
		}
	}
	return res
}
