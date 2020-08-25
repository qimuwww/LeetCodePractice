package CombinationSum

import "sort"

/*
	Given a set of candidate numbers (candidates) (without duplicates) and a target number (target), find all unique combinations in candidates where the candidate numbers sums to target.

	The same repeated number may be chosen from candidates unlimited number of times.

	Note:

	All numbers (including target) will be positive integers.
	The solution set must not contain duplicate combinations.
	Example 1:

	Input: candidates = [2,3,6,7], target = 7,
	A solution set is:
	[
	  [7],
	  [2,2,3]
	]
	Example 2:

	Input: candidates = [2,3,5], target = 8,
	A solution set is:
	[
	  [2,2,2,2],
	  [2,3,3],
	  [3,5]
	]


	Constraints:

	1 <= candidates.length <= 30
	1 <= candidates[i] <= 200
	Each element of candidate is unique.
	1 <= target <= 500
*/

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var ret [][]int
	var temp []int
	dfs(candidates, target, 0, temp, &ret)
	return ret
}

func dfs(candidates []int, target, start int, temp []int, ret *[][]int) {
	if target < 0 {
		return
	}
	if target == 0 {
		// find
		*ret = append(*ret, append([]int{}, temp...))
	}
	for i := start; i < len(candidates); i++ {
		temp = append(temp, candidates[i])
		dfs(candidates, target-candidates[i], i, temp, ret)
		temp = temp[:len(temp)-1]
	}
}

/*
func main() {
	var candidates = []int{2, 3, 5}
	var target int = 8
	ret := combinationSum(candidates, target)
	for i := range ret {
		fmt.Println(ret[i])
	}
}
//output:
[2 2 2 2]
[2 3 3]
[3 5]
*/
