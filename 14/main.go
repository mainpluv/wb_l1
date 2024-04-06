package main

import (
	"fmt"
	"reflect"
)

func getType(val interface{}) string {
	// используем reflect чтобы узнать тип элемента
	t := reflect.TypeOf(val)
	return t.String()
}

func main() {
	// создаем несколько элементов с разными типами
	intV := 20
	strV := "hi"
	boolV := true
	chV := make(chan int)
	values := []interface{}{intV, strV, boolV, chV}
	for _, v := range values {
		fmt.Println(getType(v))
	}
}
