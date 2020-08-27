package main

/*
	Find the kth largest element in an unsorted array. Note that it is the kth largest element in the sorted order, not the kth distinct element.

	Example 1:

	Input: [3,2,1,5,6,4] and k = 2
	Output: 5
	Example 2:

	Input: [3,2,3,1,2,4,5,5,6] and k = 4
	Output: 4
	Note:
	You may assume k is always valid, 1 ≤ k ≤ array's length.
*/

// 对nums数组进行降序快排
// 排序期间判断是否已经对k-1位置的数排序完成
// 随后只进行部份快排
func findKthLargest1(nums []int, k int) int {
	quickSort(nums, 0, len(nums)-1, k-1)
	return nums[k-1]
}

func quickSort(sli []int, lo, hi, k int) {
	if lo >= hi {
		return
	}
	p := partition(sli, lo, hi)
	if p == k {
		return
	} else if p > k {
		quickSort(sli, lo, p-1, k)
	} else {
		quickSort(sli, p+1, hi, k)
	}
}

func partition(sli []int, lo, hi int) int {
	povit := sli[hi]
	index := lo
	for i := lo; i < hi; i++ {
		if sli[i] > povit {
			if i != index {
				sli[i], sli[index] = sli[index], sli[i]
			}
			index++
		}
	}
	sli[index], sli[hi] = sli[hi], sli[index]
	return index
}

/*
func main() {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	ret := findKthLargest1(nums, 3)
	fmt.Println(nums)
	// output: [6 5 5 4 3 2 3 1 2]
	fmt.Println(ret)
	// output: 5
}
*/
