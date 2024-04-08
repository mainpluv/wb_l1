package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	var idx int
	// сканируем индекс элемента для удаления
	fmt.Printf("Введите индекс от 0 до %d\n", len(nums)-1)
	fmt.Scan(&idx)
	// создаем новый слайс с элементами слева и справа от удаленного числа
	res := append(nums[:idx], nums[idx+1:]...)
	fmt.Println(res)
}
