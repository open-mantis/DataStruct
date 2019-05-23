package _011_sorts

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{1, 5, 9, 6, 3, 7, 5, 10}
	fmt.Println("排序前：", arr)
	BubbleSort(arr, len(arr))
	fmt.Println("排序后：", arr)
}

func TestInsertionSort(t *testing.T) {
	arr := []int{1, 5, 9, 6, 3, 7, 5, 10}
	fmt.Println("排序前：", arr)
	InsertionSort(arr, len(arr))
	fmt.Println("排序后：", arr)
}

func TestSelectionSort(t *testing.T) {
	arr := []int{1, 5, 9, 6, 3, 7, 5, 10}
	fmt.Println("排序前：", arr)
	SelectionSort(arr, len(arr))
	fmt.Println("排序后：", arr)
}

func TestMergeSort(t *testing.T) {
	arr := []int{5, 4}
	MergeSort(arr)
	t.Log(arr)

	arr = []int{5, 4, 3, 2, 1}
	MergeSort(arr)
	t.Log(arr)
}

func createRandomArr(length int) []int {
	arr := make([]int, length, length)
	for i := 0; i < length; i++ {
		arr[i] = rand.Intn(100)
	}
	return arr
}

func TestQuickSort(t *testing.T) {
	arr := []int{81, 43, 87, 47, 59, 81}
	QuickSort(arr)
	t.Log(arr)

	arr = createRandomArr(100)
	fmt.Println(arr)
	QuickSort(arr)
	t.Log(arr)
}

func TestCountingSort(t *testing.T) {
	arr := []int{5, 4}
	CountingSort(arr, len(arr))
	t.Log(arr)

	arr = []int{0, 5, 4, 3, 2, 1}
	CountingSort(arr, len(arr))
	t.Log(arr)
}
