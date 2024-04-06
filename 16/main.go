package main

import "fmt"

func quickSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	// берем правый и левый индексы элементов массива
	l, r := 0, len(nums)-1
	// выбираем опорный элемент
	mIdx := len(nums) / 2
	m := nums[mIdx]

	// переставляем элементы так, чтобы слева все значения были меньше опорного, а справа больше
	for l <= r {
		for nums[l] < m {
			l++
		}
		for nums[r] > m {
			r--
		}
		if l <= r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}
	// рекурсивно сортируем левую и правую части
	quickSort(nums[:mIdx])
	quickSort(nums[mIdx+1:])
}

func main() {
	nums := []int{3, 8, 1, 2, 0, 45, 1, 7, 3, 8, 45, 8, 4}
	quickSort(nums)
	fmt.Println(nums)
}
