package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Введите количество клиентов:")
	var clients int

	for {
		var input string
		fmt.Scan(&input)

		value, err := strconv.Atoi(input)
		if value <= 0 || err != nil {
			fmt.Println("Введите корректное число клиентов:")
			continue
		}

		clients = value
		break
	}

	rand.Seed(time.Now().UnixNano())

	cash := make([]int, clients)
	for i := 0; i < clients; i++ {
		cash[i] = randomCash()
	}

	//cash := []int{5, 5, 5, 10, 20}

	fmt.Println("Количесвто клиентов:", clients)
	fmt.Println("Купюры клиентов:", cash)

	resultCheck := changeCheck(cash, clients)

	if resultCheck {
		fmt.Println("Киоск может предоставить каждому покупателю правильную сдачу.")
	} else {
		fmt.Println("Киоск не может предоставить каждому покупателю правильную сдачу.")
	}
}

func randomCash() int {
	cashAmount := []int{5, 10, 20}
	randomIndex := rand.Intn(len(cashAmount))
	return cashAmount[randomIndex]
}

func changeCheck(cash []int, clients int) bool {
	var cashbox []int
	lemonadePrice := 5

	for i := 0; i < clients; i++ {
		change := cash[0] - lemonadePrice

		for j := len(cashbox) - 1; j >= 0; j-- {
			if change >= cashbox[j] {
				change -= cashbox[j]
				cashbox = append(cashbox[:j], cashbox[j+1:]...)
			}
		}

		if change == 0 {
			cashbox = append(cashbox, cash[0])
			cash = cash[1:]
		} else {
			return false
		}
	}

	return true
}
