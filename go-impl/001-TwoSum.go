package main

import "fmt"
/*
	Given an array of integers, return indices of the two numbers such that they add up to a specific target.

	You may assume that each input would have exactly one solution, and you may not use the same element twice.

	Example:

	Given nums = [2, 7, 11, 15], target = 9,

	Because nums[0] + nums[1] = 2 + 7 = 9,
	return [0, 1].
*/

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	var index [2]int
	for i := 0; i < len(nums); i++ {
		numMap[nums[i]] = i
		sub := target - nums[i]
		if v, ok := numMap[sub]; ok {
			index[0], index[1] = v, i
			return index[:]
		}
	}
	return nil
}

func main()  {
	nums := []int{2, 7, 11, 15}	
	RetValue := twoSum(nums, 9)
	if RetValue != nil {
		fmt.Println(RetValue[0], RetValue[1])
	}
}