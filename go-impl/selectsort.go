package selectSort

/*
	遍历
	每次从[i,l]之间找到最小的数,然后将该数放到array[i]的位置
*/

func selectSort(array []int) {
	var tmp int
	l := len(array)
	if l <= 1 {
		return
	}
	for i := 0; i < l; i++ {
		tmp = i
		for j := i; j < l; j++ {
			if array[tmp] > array[j] {
				tmp = j
			}
		}
		array[i], array[tmp] = array[tmp], array[i]
	}
}
