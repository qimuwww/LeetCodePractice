package leetcode

/*
	Given a collection of intervals, merge all overlapping intervals.

	Example 1:

	Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
	Output: [[1,6],[8,10],[15,18]]
	Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
	Example 2:

	Input: intervals = [[1,4],[4,5]]
	Output: [[1,5]]
	Explanation: Intervals [1,4] and [4,5] are considered overlapping.
	NOTE: input types have been changed on April 15, 2019. Please reset to default code definition to get new method signature.



	Constraints:

	intervals[i][0] <= intervals[i][1]
*/

func merge(intervals [][]int) [][]int {
	if len(intervals) < 2 {
		return intervals
	}
	// 首先根据每个区间的起始位置的值进行升序排序
	miQuickSort(intervals, 0, len(intervals)-1)
	var index int
	for i := 1; i < len(intervals); i++ {
		// 比较相邻两个区间,如果前区间的结束值大于等于后区间的起始值
		// 则说明这两个区间有重合
		// 区间合并:选出这两个重合区间起始值的最小值和结束值的最大值
		// 这两个值就是合并后区间的起始值和结束值
		// 然后用这个合并后的区间再和下一个区间进行重合判断
		if intervals[index][1] >= intervals[i][0] {
			mi := min(intervals[index][0], intervals[i][0])
			ma := max(intervals[index][1], intervals[i][1])
			// 合并后的区间覆盖合并前的区间,直接复用传入的intervals的内存
			intervals[index] = []int{mi, ma}
		} else {
			index++
			// 没有重合
			intervals[index] = intervals[i]
		}
	}
	return intervals[:index+1]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func miQuickSort(sli [][]int, lo, hi int) {
	if lo >= hi {
		return
	}
	p := miPartition(sli, lo, hi)
	miQuickSort(sli, lo, p-1)
	miQuickSort(sli, p+1, hi)
}

func miPartition(sli [][]int, lo, hi int) int {
	povit := sli[hi][0]
	index := lo
	for i := lo; i < hi; i++ {
		if sli[i][0] < povit {
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
	//intervals := [][]int{{0, 3}, {0, 1}, {0, 2}, {1, 9}, {2, 5}, {10, 11}, {12, 20}, {19, 20}}
	intervals := [][]int{{1, 9}, {2, 5}, {19, 20}, {10, 11}, {12, 20}, {0, 3}, {0, 1}, {0, 2}}
	nintervals := merge(intervals)
	for i := range nintervals {
		fmt.Println(nintervals[i][0], intervals[i][1])
	}
}
*/
