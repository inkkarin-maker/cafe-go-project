package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var mood string
	for {
		fmt.Print("Какое у тебя сегодня настроение? (сияющее/грустное/злое)")
		fmt.Scan(&mood)
		if mood == "stop" {
			fmt.Println("bye-bye!")
			break
		}
		rand.Seed(time.Now().UnixNano())
		r := rand.Intn(3)
		switch mood {
		case "сияющее":
			switch r {
			case 0:
				fmt.Println("Это ураган!")
			case 1:
				fmt.Println("Это торнадо!")
			case 2:
				fmt.Println("Просто гром и молния!")
			}
			fmt.Println("Прогноз: Ослепительное солнце!")
		case "грустное":
			switch r {
			case 0:
				fmt.Println("На небе ни тучи")
			case 1:
				fmt.Println("Такой прекрасный ветерок")
			case 2:
				fmt.Println("Прекрасая радуга")
			}
			fmt.Println("Прогноз: О,нет! Кажется надвигается туча...")
		case "злое":
			switch r {
			case 0:
				fmt.Println("Это просто тучки")
			case 1:
				fmt.Println("Кажется дождик капает")
			case 2:
				fmt.Println("Ветер поднялся")
			}
			fmt.Println("О, нет. Это же настоящий ураган!!!")
		default:
		}
		time.Sleep(time.Second)
	}
}
