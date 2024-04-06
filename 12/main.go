package main

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	// создаем мапу - множество
	set := make(map[string]bool)
	for _, v := range arr {
		// если уже есть такой элемент - идем дальше
		if _, exists := set[v]; exists {
			continue
		}
		// добавляем элемент в множество
		set[v] = true
	}
	fmt.Println(set)
}
