package main

import (
	"fmt"
	"time"
)

type Order struct {
	Drink string
	Name  string
	Cost  int
}

func makeCoffee(name string, orders chan Order) {
	for o := range orders {
		fmt.Println(name, "начинает готовить ваш напиток..", o.Name, o.Drink)
		time.Sleep(time.Second * 2)
		fmt.Println("Готово!", o.Drink, "для", o.Name)
	}
}
func main() {
	orders := make(chan Order)
	go makeCoffee("Бариста №1", orders)
	go makeCoffee("Бариста №2", orders)
	orders <- Order{Drink: "Latte", Name: "Ainur", Cost: 1590}
	orders <- Order{Drink: "Cappuccino", Name: "Batyr", Cost: 1690}
	orders <- Order{Drink: "Tea", Name: "Inkar", Cost: 1290}
	close(orders)
	time.Sleep(time.Second * 10)
}
