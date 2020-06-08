package insertSort

/*
	遍历数组
	每次处理子数组array[0:i],依次逆序判断相邻的两个数的大小,如果满足条件则交换两个数
*/

func insertSort(array []int) {
	l := len(array)
	if l <= 1 {
		return
	}
	for i := 1; i < l; i++ {
		for j := i; j > 0; j-- {
			if array[j] > array[j-1] {
				array[j], array[j-1] = array[j-1], array[j]
			}
		}

	}

}
