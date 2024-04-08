package main

import "fmt"

func replace(in string) string {
	var word, res string
	for i := len(in) - 1; i >= 0; i-- {
		// если слово не закончилось, продолжаем его строить
		if in[i] != ' ' && i != 0 {
			word = string(in[i]) + word
			// если это первое слово, добавляем без пробела
		} else if in[i] == ' ' && len(word) != 0 && len(res) == 0 {
			res = word
			word = ""
			// если конец строки, заканчиваем слово и добавляем
		} else if in[i] != ' ' && i == 0 {
			word = string(in[i]) + word
			res += " " + word
			// если слово не первое, добавляем с пробелом
		} else {
			res += " " + word
			word = ""
		}
	}
	return res
}

func main() {
	var str = "snow dog sun"
	rev := replace(str)
	fmt.Println(rev)
}
