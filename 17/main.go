package main

import (
	"fmt"
	"sort"
)

func binarySearch(nums []int, trg int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		//избегаем переполнения
		mid := l + (r-l)/2
		//нашли элемент
		if nums[mid] == trg {
			return mid
		} else if nums[mid] < trg {
			l = mid + 1 // ищем справа
		} else {
			r = mid - 1 // ищем слева
		}
	}
	return -1 // элемента нет

}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 10, 24, 35, 47, 56, 68, 79, 89}
	trg := 68
	// для начала мы должны убедиться, что массив отсортирован
	sort.Ints(nums)
	idx := binarySearch(nums, trg)
	if idx != -1 {
		fmt.Println(trg, idx)
	} else {
		fmt.Println("not found")
	}
}
