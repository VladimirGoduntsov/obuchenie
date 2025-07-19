package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type User struct {
	ID   string
	Name string
}

type Account struct {
	UserID    string
	Balance   float64
	Deposits  float64
	Withdraws float64
}

func main() {

	users := readUsers("data/users.csv")
	deposits := readTransaction("data/in.csv")
	withdraws := readTransaction("data/out.csv")

	accounts := make(map[string]*Account)

	for _, user := range users {
		accounts[user.ID] = &Account{
			UserID:    user.ID,
			Balance:   0,
			Deposits:  0,
			Withdraws: 0,
		}
	}

	for userID, amount := range deposits {
		if account, exists := accounts[userID]; exists {
			account.Deposits += amount
			account.Balance += amount
		}
	}

	for userID, amount := range withdraws {
		if account, exists := accounts[userID]; exists {
			account.Withdraws += amount
			account.Balance -= amount
		}
	}
	for _, user := range users {
		if account, exists := accounts[user.ID]; exists {
			fmt.Printf("Пользователь: %s (ID: %s), Баланс: %.2f,Пополнения: %.2f, Списание: %.2f\n", user.Name, user.ID, account.Balance, account.Deposits, account.Withdraws)
		}
	}
}
func readUsers(filePath string) []User {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Ошибка при открытии файла", err)
		return []User{}
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';'

	rawData, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Ошибка при чтении данных", err)
		return []User{}

	}

	var users []User
	for _, record := range rawData {

		if len(record) == 2 {
			users = append(users, User{ID: record[0], Name: record[1]})
		}
	}
	return users
}
func readTransaction(filePath string) map[string]float64 {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Оштбка при открытии файлы", err)
		return nil
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';'

	rawData, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Ошибка при чтении файла", err)
		return nil
	}

	transactions := make(map[string]float64)
	for _, record := range rawData {

		if len(record) == 2 {
			userID := record[0]
			amount, err := strconv.ParseFloat(record[1], 64)
			if err != nil {
				fmt.Printf("Ошибка при парсинге суммы для пользователях %s: %v\n", userID, err)
				continue
			}

			transactions[userID] += amount
		}
	}
	return transactions
}
