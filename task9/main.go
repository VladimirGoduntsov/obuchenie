package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"mime"
	"net/smtp"
	"os"
	"path/filepath"
	"strings"
)

type manager struct {
	Fio   string
	Email string
	Cat   string
}

type directors struct {
	Email string
	Cat   string
}

func readDirectors(cat string) []directors {
	file, err := os.Open("directors.csv")
	if err != nil {
		fmt.Println("Файл не открывается")
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';'

	var director []directors

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		if cat == record[1] {
			director = append(director, directors{Email: record[0], Cat: record[1]})
		}
	}
	return director
}

func readManager(cat string) []manager {
	file, err := os.Open("operators.csv")
	if err != nil {
		fmt.Println("Файл не открывается")
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';'

	var operators []manager

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		if cat == record[3] {
			operators = append(operators, manager{Fio: record[0], Email: record[1], Cat: record[3]})
		}
	}

	return operators
}

func sendEmailWithCC(to []string, cc []string, subject, body, attachmentPath string) error {
	from := "your-email@yandex.ru"
	password := "your-password"

	smtpHost := "smtp.yandex.ru"
	smtpPort := "587"

	headers := make(map[string]string)
	headers["From"] = from
	headers["To"] = joinEmails(to)
	headers["Cc"] = joinEmails(cc)
	headers["Subject"] = subject
	headers["MIME-Version"] = "1.0"

	var msg bytes.Buffer
	if attachmentPath == "" {
		// Без вложения
		headers["Content-Type"] = "text/plain; charset=\"utf-8\""
		for key, value := range headers {
			msg.WriteString(fmt.Sprintf("%s: %s\r\n", key, value))
		}
		msg.WriteString("\r\n")
		msg.WriteString(body)
	} else {
		// С вложением
		headers["Content-Type"] = "multipart/mixed; boundary=\"boundary\""
		for key, value := range headers {
			msg.WriteString(fmt.Sprintf("%s: %s\r\n", key, value))
		}
		msg.WriteString("\r\n")

		// Текстовая часть
		msg.WriteString("--boundary\r\n")
		msg.WriteString("Content-Type: text/plain; charset=\"utf-8\"\r\n")
		msg.WriteString("\r\n")
		msg.WriteString(body)
		msg.WriteString("\r\n")

		// Вложение
		fileContent, err := os.ReadFile(attachmentPath)
		if err != nil {
			return fmt.Errorf("ошибка чтения файла: %w", err)
		}

		extension := filepath.Ext(attachmentPath)
		mimeType := mime.TypeByExtension(extension)
		if mimeType == "" {
			mimeType = "application/octet-stream"
		}

		filename := filepath.Base(attachmentPath)
		encodedContent := base64.StdEncoding.EncodeToString(fileContent)

		msg.WriteString("--boundary\r\n")
		msg.WriteString(fmt.Sprintf("Content-Type: %s; name=\"%s\"\r\n", mimeType, filename))
		msg.WriteString(fmt.Sprintf("Content-Disposition: attachment; filename=\"%s\"\r\n", filename))
		msg.WriteString("Content-Transfer-Encoding: base64\r\n")
		msg.WriteString("\r\n")
		msg.WriteString(encodedContent)
		msg.WriteString("\r\n--boundary--\r\n")
	}

	auth := smtp.PlainAuth("", from, password, smtpHost)
	recipients := append(to, cc...)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, recipients, msg.Bytes())
	if err != nil {
		return fmt.Errorf("ошибка отправки письма: %w", err)
	}

	return nil
}

func joinEmails(emails []string) string {
	return "\"" + strings.Join(emails, "\", \"") + "\""
}

func main() {
	// Проверка наличия аргументов командной строки
	if len(os.Args) < 2 {
			log.Fatalf("Необходимо указать категорию и (опционально) путь к файлу вложения")
	}

	// Первый аргумент — категория
	category := os.Args[1]

	// Второй аргумент (опционально) — путь к файлу вложения
	var attachmentPath string
	if len(os.Args) > 2 {
			attachmentPath = os.Args[2]
			if _, err := os.Stat(attachmentPath); os.IsNotExist(err) {
					log.Fatalf("Файл %s не найден", attachmentPath)
			}
	} else {
			attachmentPath = ""
	}

	// Чтение данных из CSV-файлов
	operators := readManager(category)
	fmt.Println("Операторы:", operators)

	directors := readDirectors(category)
	fmt.Println("Директора:", directors)

	// Чтение текста письма
	text, err := os.ReadFile("tekst.txt")
	if err != nil {
			log.Fatalf("Не удалось прочитать текст из файла: %v", err)
	}

	// Запрос темы письма
	fmt.Print("Введите тему письма: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	subject := scanner.Text()

	// Формирование списков email-адресов
	operatorEmails := make([]string, len(operators))
	for i, op := range operators {
			operatorEmails[i] = op.Email
	}

	directorEmails := make([]string, len(directors))
	for i, dir := range directors {
			directorEmails[i] = dir.Email
	}

	// Отправка письма
	err = sendEmailWithCC(operatorEmails, directorEmails, subject, string(text), attachmentPath)
	if err != nil {
			log.Printf("Ошибка отправки письма: %v", err)
	} else {
			log.Printf("Письмо успешно отправлено операторам с директорами в копии")
	}
}