package main

import (
	"fmt"
)

func quickSort(sli []int, lo, hi int) {
	if lo >= hi {
		return
	}
	p := partition(sli, lo, hi)
	quickSort(sli, lo, p-1)
	quickSort(sli, p+1, hi)
}

/*
---------------------->
3 7 8 | 5 | 2 1 9 5 | 4
------|---|---------|--
3 2 1 | 4 | 7 8 9 5 | 5

1 2 3 | 4 | 5 | 8 9 5 7
1 2 3 | 4 | 5 | 5 7 8 9

3 < 4 需要放在左半部,由于3原本就在左半部份,则3交换后还放在原地
2 < 4 2与7进行交换
1 < 4 1与8进行交换
最后基准点元素与5进行交换
此时,4之前的所有值都小于4,4之后的所有元素都大于4


基本思路:
	1. 区间分区,传入slice的一个子序列
	2. 选取一个基准点,这里选择slice的最后一个元素,简化排序
	3. 循环比较slice中的元素与基准点的元素进行比较
	4. 返回两个子序列
	5. 递归执行上述操作
*/
func partition(sli []int, lo, hi int) int {
	povit := sli[hi]
	index := lo
	for i := lo; i < hi; i++ {
		if sli[i] < povit {
			if i != index {
				sli[i], sli[index] = sli[index], sli[i]
			}
			index++
		}
	}
	sli[index], sli[hi] = sli[hi], sli[index]
	return index
}

func main() {
	s := []int{3, 7, 8, 5, 2, 1, 9, 5, 4, 1, 100, 2, 394, 485, 234, 235, 25, 25, 25, 6473465, 7, 362}
	fmt.Println(s)
	quickSort(s, 0, len(s)-1)
	fmt.Println(s)
}
