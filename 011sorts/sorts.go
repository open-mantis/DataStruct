package _011_sorts

/*
冒泡排序、插入排序、选择排序
稳定排序算法
适合小规模排序，时间复杂度O(n2)
*/

//冒泡排序，a是数组，n表示数组大小
func BubbleSort(a []int, n int) {
	if n <= 1 {
		return
	}

	for i := 0; i < n; i++ {
		flag := true
		for j := i + 1; j < n; j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
				flag = false
			}
		}

		//表示数据已经有序不需要继续循环
		if flag {
			break
		}
	}
}

// 插入排序，a表示数组，n表示数组大小
//它的工作原理是通过构建有序序列，对于未排序数据，在已排序序列中从后向前扫描，找到相应位置并插入
func InsertionSort(a []int, n int) {
	if n <= 1 {
		return
	}

	for i := 1; i < n; i++ {
		val := a[i]
		//寻找index同时移动数据
		j := i - 1
		for ; j >= 0; j-- {
			if a[j] < val {
				break
			}
			a[j+1] = a[j]
		}
		a[j+1] = val
	}
}

// 选择排序，a表示数组，n表示数组大小
//每次选择最小或者最大的元素放在已排序的队头或者对尾
func SelectionSort(a []int, n int) {
	if n <= 1 {
		return
	}
	for i := 0; i < n; i++ {
		//	查找最小的一个元素并与第i个进行交换
		min := a[i]
		var index int = 0
		for j := i + 1; j < n; j++ {
			if a[j] < min {
				min = a[j]
				index = j
			}
		}

		//交换
		if index != 0 {
			a[index] = a[i]
			a[i] = min
		}
	}
}
