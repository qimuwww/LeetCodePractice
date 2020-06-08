package bubbleSort

func bubbleSort(array []int) {
	if len(array) <= 1 {
		return
	}
	b := false
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				b = true
			}
		}
		if !b {
			break
		}
	}
}

func bubbleSort2(array []int) {
	if len(array) <= 1 {
		return
	}
	b := false
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] < array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				b = true
			}
		}
		if !b {
			break
		}
	}
}
