package _011_sorts

/*
快排
通过一趟排序将要排序的数据分割成独立的两部分，其中一部分的所有数据都比另外一部分的所有数据都要小，
然后再按此方法对这两部分数据分别进行快速排序，整个排序过程可以递归进行，以此达到整个数据变成有序序列。
*/

func QuickSort(array []int) {
	separateSort(array, 0, len(array)-1)
}

func separateSort(arr []int, start, end int) {
	if start >= end {
		return
	}
	//给出基准数字的index
	i := partition(arr, start, end)
	separateSort(arr, start, i-1)
	separateSort(arr, i, end)
}

func partition(arr []int, start, end int) int {
	//基准数字选取为 index 为end数数值
	standard := arr[end]
	i, j := start, start
	for ; j < end; j++ {
		if arr[j] <= standard {
			//交换数据位置
			if i != j {
				arr[i], arr[j] = arr[j], arr[i]
			}
			i++
		}
	}

	arr[end], arr[i] = arr[i], arr[end]
	return i
}
