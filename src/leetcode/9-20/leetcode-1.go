package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if _, ok := m[another]; ok {
			return []int{m[another], i}
		}
		m[nums[i]] = i
	}
	return nil
}

func twoSumm(nums []int, target int) []int {
	for r, s := range nums {
		for i := r + 1; i < len(nums); i++ {
			if s+nums[i] == target {
				return []int{r, i}
			}
		}
	}
	return nil
}

func main() {
	nums := []int{3, 2, 4}
	ret := twoSum(nums, 6)
	rett := twoSumm(nums, 6)
	fmt.Printf("ret=%d\n", ret)
	fmt.Printf("rett%d\n", rett)
}
