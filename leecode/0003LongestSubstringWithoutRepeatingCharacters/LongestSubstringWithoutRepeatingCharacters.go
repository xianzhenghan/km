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

// 解法三 滑动窗口-哈希桶 这个总算法很好
func lengthOfLongestSubstring2(s string) int {
	fmt.Printf("s = %+v\n", s)
	right, left, res := 0, 0, 0
	indexes := make(map[byte]int, len(s))
	for left < len(s) {
		fmt.Printf("LongestSubstring2 idx=%d;left=%d; right=%d;indexes=%+v \n",
			indexes[s[left]], left, right, indexes)
		if idx, ok := indexes[s[left]]; ok && idx >= right {
			right = idx + 1
		}
		indexes[s[left]] = left
		left++
		res = max(res, left-right)
	}
	return res
}

//自己实现
func lengthOfLongestTest(s string) int {
	left, right, res := 0, 0, 0
	indexs := make(map[byte]int, len(s))
	for right < len(s) {
		fmt.Printf("LongestSubstring2 idx=%d;left=%d; right=%d;indexes=%+v \n",
			indexs[s[left]], left, right, indexs)
		if idx, ok := indexs[s[right]]; ok && idx >= left {
			left = idx + 1
		}
		indexs[s[right]] = right
		right++
		res = max(res, right-left)
	}
	return res
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
