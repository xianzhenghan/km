package leetcode

import "fmt"

func twoSum(nums []int, target int) []int {
	fmt.Printf("twoSum target = %d,nums=%v\n", target, nums)
	mapp := make(map[int]int)
	for idx, v := range nums {
		fmt.Printf("twoSum v = %d,idx= %v,map = %v \n", v, idx, mapp)
		value, ok := mapp[target-v]
		if ok {
			return []int{idx, value}
		}
		mapp[v] = idx
	}
	return nil
}
