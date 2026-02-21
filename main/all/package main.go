package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var count int
	fmt.Print("Сколько пельменей варим?")
	fmt.Scan(&count)
	rand.Seed(time.Now().UnixNano())
	if count <= 0 {
		fmt.Print("Эй, так мы останемся голодными!")
		return
	}
	luckyNumber := rand.Intn(count) + 1
	fmt.Println("Варим...")
	time.Sleep(2 * time.Second)
	fmt.Println("Счастливым будет пельмень №...", luckyNumber)
}
