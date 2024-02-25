// FILE: N-Repeated-Element-in-Size-N-Array/N-RepeatedElementinSizeNArray.go
// AUTH: ydhmkpsryk321@gmail.com
// DES: 2024-02-25 13:44:36
package leetcode

func repeatedNTimes(A []int) int {
	kv := make(map[int]struct{})
	for _, val := range A {
		if _, ok := kv[val]; ok {
			return val
		}
		kv[val] = struct{}{}
	}
	return 0
}
