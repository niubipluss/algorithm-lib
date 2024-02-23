// FILE: Reverse-String/ReverseString.go
// AUTH: ydhmkpsryk321@gmail.com
// DES: 2024-02-23 10:13:35
package leetcode

func reverseString(s []byte) {
	for i, j := 0, len(s)-1; i < j; {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}
