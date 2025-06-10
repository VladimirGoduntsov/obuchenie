package main

import (
	"fmt"
)

func main() {

	data := "Hello World"
	data2 := Cesar(data, 3)
	fmt.Println("Кодер:", data2)

	data3 := DecodCesar(data2, 3)
	fmt.Println("Декодер:", data3)

}

func Cesar(data string, kluch int) string {
	result := make([]rune, len([]rune(data)))
	for i, v := range data {
		result[i] = v + rune(kluch)
	}
	return string(result)

}

func DecodCesar(data2 string, kluch int) string {
	result := make([]rune, len([]rune(data2)))
	for i, v := range data2 {
		result[i] = v - rune(kluch)
	}
	return string(result)

}
