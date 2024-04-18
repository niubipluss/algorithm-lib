// FILE: Find-Center-of-Star-Graph/FindCenterofStarGraph.go
// AUTH: ydhmkpsryk321@gmail.com
// DES: 2024-04-18 16:00:07
package leetcode

func findCenter(edges [][]int) int {
	if edges[0][0] == edges[1][0] || edges[0][0] == edges[1][1] {
		return edges[0][0]
	}
	return edges[0][1]
}
