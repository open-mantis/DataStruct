package _011_sorts

/*
归并排序
核心思路采用递归分区治理，分区有序，然后在进行合并排序得出最后结果
需要占用额外空间最大为队列本身大小
稳定的排序算法
时间复杂度 nlog(n)
*/

func MergeSort(array []int) {
	arrLen := len(array)
	if arrLen <= 1 {
		return
	}
	mergeSort(array, 0, arrLen-1)
}

func mergeSort(array []int, start, end int) {
	if start >= end {
		return
	}
	mid := (start + end) / 2
	mergeSort(array, start, mid)
	mergeSort(array, mid+1, end)
	merge(array, start, mid, end)

}

func merge(array []int, start, mid, end int) {
	//用于存放排序的临时数组
	tmpArray := make([]int, end-start+1)
	i := start
	j := mid + 1
	index := 0

	//按顺序将节点放入临时数组
	for ; i <= mid && j <= end; index++ {
		if array[i] < array[j] {
			tmpArray[index] = array[i]
			i++
		} else {
			tmpArray[index] = array[j]
			j++
		}
	}

	//剩余的节点更新至tmp array
	for ; i <= mid; i++ {
		tmpArray[index] = array[i]
		index++
	}

	for ; j <= end; j++ {
		tmpArray[index] = array[j]
		index++
	}
	copy(array[start:end+1], tmpArray)
	return
}
