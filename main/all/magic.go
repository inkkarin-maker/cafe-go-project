package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main1() {
	var name string
	fmt.Print("Как вас зовут?")
	fmt.Scan(&name)
	predictions := []string{"Сегодня будет чудесный день!", "Будь осторожен!"}
	rand.Seed(time.Now().UnixNano())
	var i = rand.Intn(len(predictions))
	fmt.Println(name, ",твой прогноз на сегодня:", predictions[i])
}
