package leetcode

import "fmt"

func LongestSubstringWithoutRepeatingCharacters(s string) int {
	fmt.Printf("s = %+v\n", s)
	if len(s) == 0 {
		return 0
	}
	var freq [127]int
	result, left, right := 0, 0, -1
	for left < len(s) {
		fmt.Printf("LongestSubstring  left=%d;	right=%d  ", left, right)
		print_nozero(freq)
		if right+1 < len(s) && freq[s[right+1]] == 0 {
			freq[s[right+1]]++
			right++
		} else {
			freq[s[left]]--
			left++
		}
		result = max(result, right-left+1)
	}
	return result
}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func print_nozero(nums [127]int) {
	fmt.Printf(" freq :")
	for idx, v := range nums {
		if v > 0 {
			fmt.Printf("%s => %d	", string(rune(idx)), v)
		}
	}
	fmt.Printf("\n")
}
