package main

import (
	"fmt"
	"strings"
)

func main() {

	text := `макароны — это палочка-выручалочка, если вам нужно быстро приготовить что-то вкусное и сытное. Но и у этого, казалось бы, простого гарнира есть свои секреты. Сколько времени варить макароны как сделать так, чтобы они не слипались и нужно ли их промывать — эти вопросы часто возникают не только у новичков, но и уже опытных кулинаров. Подробно разберем тонкости, как правильно готовить блюдо в кастрюле, мультиварке и микроволновке`

	words := strings.Fields(text)

	hraniliche := make(map[string]int)
	for _, w := range words {
		hraniliche[w]++

	}
	//fmt.Println(hraniliche)

	var repeatedWord string // Хранение часто повторяемого слова
	numberWords := 0        // Хранение количеста - сколько раз это слово встречается

	for wor, num := range hraniliche {
		if num > numberWords {
			repeatedWord = wor
			numberWords = num

		}

	}
	fmt.Println("повторяемое слово:", repeatedWord, "\n кол повторений:", numberWords)
}
