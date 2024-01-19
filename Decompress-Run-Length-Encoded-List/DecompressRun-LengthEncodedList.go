// FILE: Decompress-Run-Length-Encoded-List/DecompressRun-LengthEncodedList.go
// AUTH: ydhmkpsryk321@gmail.com
// DES: 2024-01-19 11:00:20
package leetcode

func decompressRLElist(nums []int) []int {
	res := []int{}
	for i := 0; i < len(nums); i += 2 {
		for j := 0; j < nums[i]; j++ {
			res = append(res, nums[i+1])
		}
	}
	return res
}
