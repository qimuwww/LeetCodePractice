package twoSum

//import "fmt"
/*
	Given an array of integers, return indices of the two numbers such that they add up to a specific target.

	You may assume that each input would have exactly one solution, and you may not use the same element twice.

	Example:

	Given nums = [2, 7, 11, 15], target = 9,

	Because nums[0] + nums[1] = 2 + 7 = 9,
	return [0, 1].
*/

func twoSum(nums []int, target int) []int {
	var (
		numMap = make(map[int]int)
		index  int
		ok     bool
		value  [2]int
	)
	for i := 0; i < len(nums); i++ {
		if index, ok = numMap[nums[i]]; ok {
			if nums[i]*2 == target {
				value[0], value[1] = index, i
				return value[:]
			}
		}
		numMap[nums[i]] = i
		if index, ok = numMap[target-nums[i]]; ok {
			if index != i {
				value[0], value[1] = index, i
				return value[:]
			}

		}
	}
	return nil
}

/*
func main()  {
	nums := []int{3, 3}
	RetValue := twoSum(nums, 6)
	if RetValue != nil {
		fmt.Println(RetValue[0], RetValue[1])
	}
}
*/
