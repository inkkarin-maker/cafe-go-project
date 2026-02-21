package main

import (
	"errors"
	"fmt"
)

type Robot struct {
	Name  string
	Model string
	Power int
}

func (r Robot) SayHello() {
	fmt.Printf("Привет! Я робот %s, моя мощность %d ватт!\n", r.Name, r.Power)
}

func (r *Robot) Charge() {
	r.Power += 1000
	fmt.Printf("%s заряжен! Новая мощность: %d\n", r.Name, r.Power)
}

type Speaker interface {
	SayHello()
}
type Human struct {
	Name string
}

func (h Human) SayHello() {
	fmt.Printf("Привет! Я человек, меня зовут %s. Я не робот (честно)!\n", h.Name)
}

type Dog struct {
	Name string
}

func (d Dog) SayHello() {
	fmt.Printf("Гав! Я пес %s, я просто пришел за печеньками!\n", d.Name)
}
func CheckSafety(power int) error {
	if power > 4000 {
		return errors.New("КРИТИЧЕСКАЯ МОЩНОСТЬ! Робот может взорваться!")
	}
	return nil
}
func main() {
	report := make(chan int)
	myRobot := Robot{Name: "Pi", Model: "xy", Power: 2500}
	myRegistry := make(map[string]Robot)
	myRegistry["Pi"] = myRobot
	bot, ok := myRegistry["Unknown"]
	if ok {
		fmt.Println("Робот найден:", bot.Name)
	} else {
		fmt.Println("Ошибка: Такого робота в базе нет!")
	}
	ivan := Human{Name: "Ivan"}
	party := []Speaker{myRobot, ivan}
	fmt.Println("--- Вечеринка начинается! ---")
	for _, guest := range party {
		guest.SayHello()
	}
	army := []*Robot{
		&myRobot,
		&Robot{Name: "Sparky", Model: "xz", Power: 3000},
	}
	myRobot.Charge()
	for _, bot := range army {
		go func(b *Robot) {
			b.Charge()
			report <- b.Power
		}(bot)
	}
	myRobot.SayHello()
	fmt.Println("My robot's power is", myRobot.Power, "v")
	sharik := Dog{Name: "Шарик"}
	party = append(party, sharik)
	result1 := <-report
	result2 := <-report
	fmt.Printf("Отчет получен! Мощности в канале: %d и %d\n", result1, result2)

	err := CheckSafety(myRobot.Power)
	if err != nil {
		fmt.Println("ВНИМАНИЕ:", err)
	} else {
		fmt.Println("Система стабильна. Мощность в пределах нормы.")
	}
}
