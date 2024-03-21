// FILE: Assign-Cookies/AssignCookies.go
// AUTH: ydhmkpsryk321@gmail.com
// DES: 2024-03-21 16:34:59
package leetcode

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	gi, si, res := 0, 0, 0
	for gi < len(g) && si < len(s) {
		if s[si] >= g[gi] {
			res++
			si++
			gi++
		} else {
			si++
		}
	}
	return res
}
