package main

import (
	"bufio"
	"fmt"
	"os"
	//"strings"
	"unicode"
)

func main() {

	filePath := os.Args[1] // содержит первый аргумент командной строки - путь к файлу. os.Args – это срез (slice) строк, который предоставляет доступ к аргументам командной строки, переданным при запуске программы. Из пакета "os"

	//Открываем файл
	file, err := os.Open(filePath) // Функция os.Open(filePath) открывает файл, указанный в переменной filePath, в режиме только для чтения.
	if err != nil {
		fmt.Printf("Ошибка при открытии файла: %v\n", err)
		return
	}
	defer file.Close() // это ключевое слово в Go, которое откладывает выполнение указанной операции до завершения текущей функции. Гарантирует, что файл будет закрыт независимо от того, какой путь выполнения программы был выбран. Это помогает избежать утечек ресурсов.

	// Инициализация переменных для анализа
	var totalChars, // общее количество символов
		totalWords, // общее количество слов
		totalLines, // общее количество строк
		totalSentences int // общее количество предложений
	scanner := bufio.NewScanner(file)

	// Чтение файла построчно
	for scanner.Scan() { // это метод из пакета bufio, который используется для построчного чтения данных из файла.
		line := scanner.Text()                     // возвращает содержимое текущей строки, которую scanner прочитал в текущей итерации цикла. Сохраняет в переменную line
		totalLines++                               // Переменная totalLines используется для подсчета общего количества строк в тексте.
		totalChars += countNonSpaceChars(line) + 1 // Вычисляет длину текущей строки (len(line)) и добавляет её к переменной totalChars.  +1 добавляется для учета символа переноса строки (\n), который не входит в результат scanner.Text().
		totalWords += countWords(line)             // Вызывает функцию countWords(line), которая, подсчитывает количество слов в текущей строке. Функция countWords она разделяет строку на слова.
		totalSentences += countSentences(line)     // Вызывает функцию countSentences(line), которая, подсчитывает количество предложений в текущей строке. Функция countSentences она анализирует строку на наличие знаков препинания (точек, восклицательных или вопросительных знаков), которые обозначают конец предложения.
	}

	// Вывод результатов
	fmt.Printf("Общее количество символов: %d\n", totalChars)
	fmt.Printf("Общее количество слов: %d\n", totalWords)
	fmt.Printf("Общее количество строк: %d\n", totalLines)
	fmt.Printf("Общее количество предложений: %d\n", totalSentences)
}

// Функция для подсчёта символов без пробелов
func countNonSpaceChars(line string) int {
	count := 0
	for _, char := range line {
		if !unicode.IsSpace(char) {
			count++
		}
	}
	return count
}

// Функция для подсчёта слов в строке
// func countWords(line string) int {
// words := strings.FieldsFunc(line, func(r rune) bool {
// return !unicode.IsLetter(r) && !unicode.IsDigit(r)
// })
// return len(words)
func countWords(line string) int {
	var count int   // Счётчик слов, изначально равен 0.
	var inWord bool // Логическая переменная, которая указывает, находимся ли мы внутри слова (true) или нет (false). Изначально она равна false

	for _, r := range line {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			// Если текущий символ является буквой или цифрой и мы ещё не в слове,
			// начинаем новое слово.
			if !inWord {
				count++
				inWord = true
			}
		} else {
			// Если текущий символ не является буквой или цифрой, выходим из слова.
			inWord = false
		}
	}

	return count // После завершения цикла возвращается общее количество слов
}

// Функция для подсчёта предложений в строке
//func countSentences(line string) int {
//sentences := strings.FieldsFunc(line, func(r rune) bool {
//return r == '.' || r == '!' || r == '?'
//})
//return len(sentences)

func countSentences(line string) int {
	var count int // счётчик предложений, иpначально равен 0.
	for _, r := range line {
		if r == '.' || r == '!' || r == '?' { //|| r == ',' || r == ':'
			count++
		}
	}
	return count
}
