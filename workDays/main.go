package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Дни недели.")

	answer := false

	for !answer {
		fmt.Println("Введите будний день недели: пн, вт, ср, чт, пт:")
		var day string
		fmt.Scan(&day)
		day = strings.ToLower(day)

		switch day {
		case "пн":
			fmt.Println("понедельник")
			fallthrough
		case "вт":
			fmt.Println("вторник")
			fallthrough
		case "ср":
			fmt.Println("среда")
			fallthrough
		case "чт":
			fmt.Println("четверг")
			fallthrough
		case "пт":
			fmt.Println("пятница")
			answer = true
		default:
			fmt.Println("Ввод некорректен. Попробуйте снова.")
		}
	}
}
