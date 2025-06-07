package main

import (
	"fmt"
	
)

func main() {
	data := "Hello World"
	//data2 := []rune(data)
	//fmt.Println(data2)
	fmt.Println(data)
	
	for _, v := range data {
		fmt.Println(string(v+3))
	

	}
	
}