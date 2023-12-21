// FILE: Jump-Game-III/JumpGameIII.go
// AUTH: ydhmkpsryk321@gmail.com
// DES: 2023-12-21 14:24:52
package leetcode

func canReach(arr []int, start int) bool {
	if start >= 0 && start < len(arr) && arr[start] < len(arr) {
		jump := arr[start]
		arr[start] += len(arr)
		return jump == 0 || canReach(arr, start+jump) || canReach(arr, start-jump)
	}
	return false
}
