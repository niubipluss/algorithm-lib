// FILE: LeetCode-Go-master/leetcode/Jump-Game-II/JumpGameII.go
// AUTH: ydhmkpsryk321@gmail.com
// DES: 2024-04-07 10:06:33
package leetcode

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	needChoose, canReach, step := 0, 0, 0
	for i, x := range nums {
		if i+x > canReach {
			canReach = i + x
			if canReach >= len(nums)-1 {
				return step + 1
			}
		}
		if i == needChoose {
			needChoose = canReach
			step++
		}
	}
	return step
}
