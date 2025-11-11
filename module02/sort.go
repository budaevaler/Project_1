package main

import "fmt"

func main() {
	arr := []int{1111, 2025, 70, 3, 2014, 90, 760, 430}
	quick_sort(arr)
	fmt.Println(arr)
}

func partition(arr []int, min_i int, max_i int) int {
	pivot := min_i - 1

	for i := min_i; i < max_i; i++ {
		if arr[i] <= arr[max_i] {
			pivot++
			arr[i], arr[pivot] = arr[pivot], arr[i]
		}
	}

	pivot++
	arr[pivot], arr[max_i] = arr[max_i], arr[pivot]
	return pivot
}

func quick_sort_impl(arr []int, min_i int, max_i int) {
	if min_i < max_i {
		q := partition(arr, min_i, max_i)
		quick_sort_impl(arr, min_i, q-1)
		quick_sort_impl(arr, q+1, max_i)
	}
}

func quick_sort(arr []int) {
	if len(arr) != 0 {
		quick_sort_impl(arr, 0, len(arr)-1)
	}
}
