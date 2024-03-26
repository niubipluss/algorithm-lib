// FILE: Valid-Boomerang/ValidBoomerang.go
// AUTH: ydhmkpsryk321@gmail.com
// DES: 2024-03-26 11:41:19
package leetcode

func isBoomerang(points [][]int) bool {
	return (points[0][0]-points[1][0])*(points[0][1]-points[2][1]) != (points[0][0]-points[2][0])*(points[0][1]-points[1][1])
}
