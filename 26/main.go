package main

import (
	"fmt"
	"unicode"
)

func Check(str string) bool {
	// создаем мапу для проверки на уникальность
	mp := make(map[rune]bool)
	for i := 0; i < len(str); i++ {
		// все символы приводим к нижнему регистру
		letter := unicode.ToLower(rune(str[i]))
		// проверяем, существует ли уже такой символ в мапе
		if _, exists := mp[letter]; !exists {
			mp[letter] = true
		} else {
			return false
		}
	}
	return true
}

func main() {
	strs := []string{"abcd", "abCdefAaf", "aabcd"}
	for _, v := range strs {
		flag := Check(v)
		fmt.Println(v, flag)
	}
}
