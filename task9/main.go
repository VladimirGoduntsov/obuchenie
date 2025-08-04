package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/smtp"
	"os"
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

func sendEmailWithCC(to []string, cc []string, subject, body string, attachmentPath string) error {
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
	headers["Content-Type"] = "multipart/mixed; boundary=\"boundary\""

	var msg bytes.Buffer
	for key, value := range headers {
		msg.WriteString(fmt.Sprintf("%s: %s\r\n", key, value))
	}
	msg.WriteString("\r\n")

	msg.WriteString("--boundary\r\n")
	msg.WriteString("Content-Type: text/plain; charset=\"utf-8\"\r\n")
	msg.WriteString("\r\n")
	msg.WriteString(body)
	msg.WriteString("\r\n")

	msg.WriteString("--boundary\r\n")
	msg.WriteString("Content-Type: application/pdf; name=\"invoice.pdf\"\r\n")
	msg.WriteString("Content-Disposition: attachment; filename=\"invoice.pdf\"\r\n")
	msg.WriteString("Content-Transfer-Encoding: base64\r\n")
	msg.WriteString("\r\n")

	msg.WriteString("\r\n--boundary--\r\n")

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

	category := "море"
	operators := readManager(category)
	fmt.Println("Операторы море:", operators)

	directors := readDirectors(category)
	fmt.Println("Директора море:", directors)

	text, err := os.ReadFile("tekst.txt")
	if err != nil {
		log.Fatalf("Не удалось прочитать текст из файла: %v", err)
	}

	fmt.Print("Введите тему письма: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	subject := scanner.Text()

	operatorEmails := make([]string, len(operators))
	for i, op := range operators {
		operatorEmails[i] = op.Email
	}

	directorEmails := make([]string, len(directors))
	for i, dir := range directors {
		directorEmails[i] = dir.Email
	}

	err = sendEmailWithCC(operatorEmails, directorEmails, subject, string(text), "invoice.pdf")
	if err != nil {
		log.Printf("Ошибка отправки письма: %v", err)
	} else {
		log.Printf("Письмо успешно отправлено операторам с директорами в копии")
	}
}
