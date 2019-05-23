package _15binarysearch

/*
arr为有序数组
*/

//非递归方式
func BinarySearch(arr []int, v int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}

	start := 0
	end := len(arr) - 1

	/* mid可能出现溢出问题
	mid := (start + end) / 2
	调整如下:
	mid := (end-start)/2 + start
	优化除2运算
	mid := (end-start) >> 1 + start
	*/

	for start <= end {
		mid := (end-start)>>1 + start
		if arr[mid] == v {
			return mid
		} else if arr[mid] > v {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return -1
}

//二分查找递归方式
func BinarySearchRecursive(arr []int, v int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}

	return bsSearch(arr, 0, n-1, v)
}

func bsSearch(arr []int, start, end, val int) int {
	if start > end {
		return -1
	}
	mid := (end-start)>>1 + start
	if arr[mid] == val {
		return mid
	} else if arr[mid] > val {
		return bsSearch(arr, start, mid-1, val)
	} else {
		return bsSearch(arr, mid+1, end, val)
	}
}

//查找第一个等于给定值的元素
//关键要处理存在的重复元素
func BinarySearchFirst(arr []int, v int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}
	start := 0
	end := len(arr) - 1
	for start <= end {
		mid := (end-start)>>1 + start
		if arr[mid] == v {
			if mid == 0 || arr[mid-1] != v {
				return mid
			} else {
				end = mid - 1
			}
		} else if arr[mid] > v {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return -1
}

//查找最后一个值等于给定值的元素
func BinarySearchLast(arr []int, v int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}
	start := 0
	end := len(arr) - 1
	for start <= end {
		mid := (end-start)>>1 + start
		if arr[mid] == v {
			if mid == 0 || arr[mid+1] != v {
				return mid
			} else {
				start = mid + 1
			}
		} else if arr[mid] > v {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return -1
}

//查找第一个大于等于给定值的元素
func BinarySearchFirstGT(arr []int, v int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}
	start := 0
	end := len(arr) - 1
	for start <= end {
		mid := (end-start)>>1 + start
		if arr[mid] == v {
			if mid != n-1 && arr[mid+1] > v {
				return mid + 1
			} else {
				start = mid + 1
			}
		} else if arr[mid] > v {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return -1
}

//查找最后一个小于等于给定值的元素
func BinarySearchLastLT(arr []int, v int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}
	start := 0
	end := len(arr) - 1
	for start <= end {
		mid := (end-start)>>1 + start
		if arr[mid] == v {
			if mid != 0 && arr[mid-1] < v {
				return mid - 1
			} else {
				end = mid - 1
			}
		} else if arr[mid] > v {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return -1
}
