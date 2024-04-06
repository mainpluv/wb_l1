package main

import "fmt"

func intersection(s1, s2 map[int]bool) map[int]bool {
	res := make(map[int]bool)
	//ищем пересечения
	for i := range s1 {
		if s2[i] {
			res[i] = true
		}
	}
	return res
}

func main() {
	// создаем 2 множества
	s1 := map[int]bool{1: true, 2: true, 12: true, 14: true}
	s2 := map[int]bool{2: true, 3: true, 10: true, 12: true}
	// создаем подмножество пересечения
	intersect := intersection(s1, s2)
	fmt.Println(intersect)
}
