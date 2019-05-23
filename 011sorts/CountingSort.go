package _011_sorts

import (
	"fmt"
	"math"
)

func CountingSort(arr []int, n int) {
	if n <= 1 {
		return
	}
	//	获取最大值,数据均为正整数
	max := math.MinInt32
	for _, val := range arr {
		if val >= max {
			max = val
		}
	}
	fmt.Println(arr, max)
	//创建桶
	count := make([]int, max+1)

	//进行计数
	for _, val := range arr {
		count[val]++
	}

	//创建新数组
	newArr := make([]int, n)

	//计算每个原数的末尾index
	//数据存储为对应新数据val的末尾index值-1
	for i := 1; i <= max; i++ {
		count[i] += count[i-1]
	}

	//放入数据至新数组
	for i := range arr {
		index := count[arr[i]] - 1
		newArr[index] = arr[i]
		count[arr[i]]--
	}
	copy(arr, newArr)
}
