// FILE: Poor-Pigs/PoorPigs.go
// AUTH: ydhmkpsryk321@gmail.com
// DES: 2024-01-15 16:58:32
package leetcode

import "math"

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	base := minutesToTest/minutesToDie + 1
	return int(math.Ceil(math.Log10(float64(buckets)) / math.Log10(float64(base))))
}
