package main

import (
	"fmt"
	
)

func main() {

	data := "Vova"
	data2 := Cesar(data, 3)
	fmt.Println(data2)
	
}

func Cesar(data string, kluch int) string {
	result := make([]rune, len([]rune(data)))
	for i, v := range data {
	result[i] = v+rune(kluch)
	}
	return string(result)
}