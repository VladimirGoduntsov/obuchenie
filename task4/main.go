package main

import (
	"fmt"
)

func main() {

	data := []string{

		"12345678910",
		"12345678",
		"123456",
		"12345",
		"1234",
		"mnhukdsertp",
		"fjgg;b",
		"kkfoflfmvu",
		"12345678910",
		"12345678",
		"123456",
		"12345",
		"1234",
		"mnhukdsertp",
		"fjgg;b",
		"kkfoflfmvu",
		"12345678910",
		"12345678",
		"123456",
		"12345",
		"1234",
		"mnhukdsertp",
		"fjgg;b",
		"kkfoflfmvu",
		"12345678910",
		"12345678",
		"123456",
		"12345",
		"1234",
		"mnhukdsertp",
		"fjgg;b",
		"kkfoflfmvu",
		"12345678910",
		"12345678",
		"123456",
		"12345",
		"1234",
		"mnhukdsertp",
		"fjgg;b",
		"kkfoflfmvu",
		"12345678910",
		"12345678",
		"123456",
		"12345",
		"1234",
		"mnhukdsertp",
		"fjgg;b",
		"kkfoflfmvu",
		"12345678910",
		"12345678",
		"123456",
		"12345",
		"1234",
		"mnhukdsertp",
		"fjgg;b",
		"kkfoflfmvu",
		"12345678910",
		"12345678",
		"123456",
		"12345",
		"1234",
		"mnhukdsertp",
		"fjgg;b",
		"kkfoflfmvu",
		"12345678910",
		"12345678",
		"123456",
		"12345",
		"1234",
		"mnhukdsertp",
		"fjgg;b",
		"kkfoflfmvu",
		"12345678910",
		"12345678",
		"123456",
		"12345",
		"1234",
		"mnhukdsertp",
		"fjgg;b",
		"kkfoflfmvu",
		"12345678910",
		"12345678",
		"123456",
		"12345",
		"1234",
		"mnhukdsertp",
		"fjgg;b",
		"kkfoflfmvu",
		"12345678910",
		"12345678",
		"123456",
		"12345",
		"1234",
		"mnhukdsertp",
		"fjgg;b",
		"kkfoflfmvu",
		"1234",
		"mnhukdsertp",
		"fjgg;b",
		"kkfoflfmvu",
	}

	maxData2 := 0  // var maxData2 int (переменная для хранения макс длинна)
	maxData3 := "" // var maxData3 string (переменная для хранения макс строки)

	for _, v := range data {
		if len(v) > maxData2 {
			maxData2 = len(v)
			maxData3 = v
		}
	}
	fmt.Println("Максимальная длинна:", maxData2)
	fmt.Println("Строка с макс длинной:", maxData3)

	var maxData4 []string
	for _, a := range data {
		if len(a) == maxData2 {
			maxData4 = append(maxData4, a)
		}
	}
	fmt.Println("Все строки с макс длинной:", maxData4)
}
