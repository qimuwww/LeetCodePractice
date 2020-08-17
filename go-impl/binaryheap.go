package main

import "fmt"

/*
	二叉堆,一种完全二叉树,即除最后一层外的其余层都是满的，并且最后一层要么是满的，要么在右边缺少连续若干节点
	这种数据结构可以用数组来表示,从数组下标为1的位置开始存储数据
	通过根节点下标可以计算出左孩子和右孩子的下标,相反也可以通过孩子的下标算出父节点的下标

	实现:优先级队列,topN等等...
*/

// 这里只实现大顶堆,小顶堆原理类似

func parent(child int) int {
	return child / 2
}

// 左孩子的索引
func left(root int) int {
	return root * 2
}

// 右孩子的索引
func right(root int) int {
	return root*2 + 1
}

type BinaryHeap []int

func (bh *BinaryHeap) less(i, j int) bool {
	return (*bh)[i] < (*bh)[j]
}

// 上浮操作 如果孩子节点的值大于父节点的值,则交换二者
func (bh *BinaryHeap) swim(k int) {
	for {
		if (k > 1) && bh.less(parent(k), k) {
			(*bh)[parent(k)], (*bh)[k] = (*bh)[k], (*bh)[parent(k)]
			k = parent(k)
		} else {
			break
		}
	}
}

// 下沉操作,如果父节点小于两个子节点中最大的,则交换二者
func (bh *BinaryHeap) sink(k int) {
	for {
		if left(k) <= len(*bh) {
			// 先假设左边节点较大
			older := left(k)
			// 如果右边节点存在，比一下大小
			if right(k) <= len(*bh) && bh.less(older, right(k)) {
				older = right(k)
			}
			// 结点 k 比俩孩子都大，就不必下沉了
			if bh.less(older, k) {
				break
			}
			// 否则，不符合最大堆的结构，下沉 k 结点
			(*bh)[k], (*bh)[older] = (*bh)[older], (*bh)[k]
			k = older
		} else {
			break
		}
	}
}

// 向大顶堆中插入数据,先将数据插入到数组末尾,然后对该数据进行上浮操作
func (bh *BinaryHeap) Insert(key int) {
	*bh = append(*bh, key)
	bh.swim(len(*bh) - 1)
}

// 删除大顶堆堆顶数据
// 交换堆顶与堆尾的数据,删除堆尾数据,对堆顶数据进行下沉操作
func (bh *BinaryHeap) DeleteMax() int {
	key := (*bh)[1]
	(*bh)[1], (*bh)[len(*bh)-1] = (*bh)[len(*bh)-1], (*bh)[1]
	*bh = (*bh)[:len(*bh)-1]
	bh.sink(1)
	return key
}

func main() {
	b := make(BinaryHeap, 1, 20)
	b.Insert(2)
	b.Insert(1)
	b.Insert(9)
	b.Insert(6)
	b.Insert(7)
	b.Insert(8)
	b.Insert(4)
	b.Insert(5)
	b.Insert(3)
	for i := range b {
		fmt.Print(b[i])
	}
	fmt.Println()
	b.DeleteMax()
	for i := range b {
		fmt.Print(b[i])
	}
	fmt.Println()
	b.DeleteMax()
	for i := range b {
		fmt.Print(b[i])
	}
	fmt.Println()
}
