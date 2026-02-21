package main

import "fmt"

type Spaceship struct {
	Name     string
	Fuel     float64
	Time     int
	Distance int
	MaxFuel  float64
}

func (s *Spaceship) Refuel(amount float64) {
	if s.Fuel+amount > s.MaxFuel {
		s.Fuel = s.MaxFuel
		fmt.Println("Бак полон! Лишнее топливо слито.")
	} else {
		s.Fuel += amount
		fmt.Printf("Заправлено %.2f. Теперь в баке: %.2f\n", amount, s.Fuel)
	}
}

func (s *Spaceship) Fly(dist float64) {
	if s.Fuel < dist {
		fmt.Printf("Ошибка! Недостаточно топлива для полета на %.2f\n", dist)
	} else {
		s.Fuel -= dist
		s.Distance += int(dist)
		fmt.Printf("%s пролетел %.2f км. Топлива осталось: %.2f\n", s.Name, dist, s.Fuel)
	}
}
func main5() {
	myShip := Spaceship{
		Name:     "Arman",
		Fuel:     50000.0,
		Time:     24,
		MaxFuel:  100000.0,
		Distance: 800000,
	}
	var userDist float64
	fmt.Println("Введите дистанцию для полета:")
	fmt.Scan(&userDist)
	myShip.Fly(userDist)

	var fuelAmount float64
	fmt.Println("Сколько литров заправить?")
	fmt.Scan(&fuelAmount)
	myShip.Refuel(fuelAmount)

	fmt.Printf("\nИтог: Путь %d км, Топливо %.2f\n", myShip.Distance, myShip.Fuel)
}
