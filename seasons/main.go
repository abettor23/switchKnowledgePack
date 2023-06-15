package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Времена года.")

answer:
	for {
		fmt.Println("Введите месяц:")
		var season string
		fmt.Scan(&season)
		season = strings.ToLower(season)

		switch season {
		case "декабрь", "январь", "февраль":
			fmt.Println("Время года - зима")
			break answer
		case "март", "апрель", "май":
			fmt.Println("Время года - весна")
			break answer
		case "июнь", "июль", "август":
			fmt.Println("Время года - лето")
			break answer
		case "сентябрь", "октябрь", "ноябрь":
			fmt.Println("Время года - осень")
			break answer
		default:
			fmt.Println("Ввод некорректен. Попробуйте снова.")
		}
	}
}
