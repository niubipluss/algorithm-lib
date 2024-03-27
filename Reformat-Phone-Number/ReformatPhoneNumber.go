// FILE: Reformat-Phone-Number/ReformatPhoneNumber.go
// AUTH: ydhmkpsryk321@gmail.com
// DES: 2024-03-27 15:54:52
package leetcode

import (
	"strings"
)

func reformatNumber(number string) string {
	parts, nums := []string{}, []rune{}
	for _, r := range number {
		if r != '-' && r != ' ' {
			nums = append(nums, r)
		}
	}
	threeDigits, twoDigits := len(nums)/3, 0
	switch len(nums) % 3 {
	case 1:
		threeDigits--
		twoDigits = 2
	case 2:
		twoDigits = 1
	default:
		twoDigits = 0
	}
	for i := 0; i < threeDigits; i++ {
		s := ""
		s += string(nums[0:3])
		nums = nums[3:]
		parts = append(parts, s)
	}
	for i := 0; i < twoDigits; i++ {
		s := ""
		s += string(nums[0:2])
		nums = nums[2:]
		parts = append(parts, s)
	}
	return strings.Join(parts, "-")
}
