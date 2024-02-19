// FILE: Minimum-Moves-to-Equal-Array-Elements-II/MinimumMovestoEqualArrayElementsII.go
// AUTH: ydhmkpsryk321@gmail.com
// DES: 2024-02-19 11:28:16
package leetcode

import (
	"math"
	"sort"
)

func minMoves2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	moves, mid := 0, len(nums)/2
	sort.Ints(nums)
	for i := range nums {
		if i == mid {
			continue
		}
		moves += int(math.Abs(float64(nums[mid] - nums[i])))
	}
	return moves
}
