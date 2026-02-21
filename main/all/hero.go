package main

import "fmt"

type Hero struct {
	Name      string
	Health    int
	Power     float64
	Inventory []string
}

func (h *Hero) TakeDamage(damage int) {
	h.Health -= damage
	fmt.Printf("%s получил удар! Осталось здоровья: %d\n", h.Name, h.Health)
}
func (h *Hero) IsAlive() bool {
	if h.Health > 0 {
		return true
	} else {
		return false
	}
}

func (h *Hero) Heal(amount int) {
	h.Health += amount
	fmt.Printf("%s исцелился на %d. Текущее здоровье: %d\n", h.Name, amount, h.Health)
}

func main4() {
	myHero := Hero{
		Name:      "Batyr",
		Health:    35,
		Power:     70.0,
		Inventory: []string{"Old Shield"},
	}

	fmt.Println("Имя героя:", myHero.Name)

	for myHero.Power <= 100 {
		fmt.Printf("Текущая сила: %.2f. %s тренируется...\n", myHero.Power, myHero.Name)

		var newTraining int
		fmt.Println("Введите очки тренировки:")
		fmt.Scan(&newTraining)

		myHero.Power += float64(newTraining)
	}

	fmt.Println("\n--- Тренировка окончена! Выберите награду ---")
	fmt.Println("1-Magic Cape (+10 Power)")
	fmt.Println("2-Fire Boots (+15 Power)")
	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		myHero.Inventory = append(myHero.Inventory, "Magic Cape")
		myHero.Power += 10
		fmt.Println("Вы надели плащ!")
	case 2:
		myHero.Inventory = append(myHero.Inventory, "Fire Boots")
		myHero.Power += 15
		fmt.Println("Вы надели сапоги!")
	default:
		myHero.Inventory = append(myHero.Inventory, "Rusty Spoon")
		fmt.Println("Вы нашли ложку...")
	}

	fmt.Println("\nВНЕЗАПНАЯ ЛОВУШКА!")
	myHero.TakeDamage(15)

	fmt.Println("Вы нашли лечебное зелье!")
	myHero.Heal(10)

	bossHP := 500.0
	for bossHP > 0 {
		bossHP -= myHero.Power
		fmt.Printf("Босс получил удар. Осталось HP: %.2f\n", bossHP)

		myHero.TakeDamage(20)

		if myHero.IsAlive() == false {
			fmt.Println("Батыр пал в бою... Гейм овер.")
		}
	}
	if myHero.IsAlive() {
		fmt.Println("Поздравляем! Вы победили босса и выжили!")
	}

	fmt.Printf("\nГерой: %s | Финальная сила: %.2f\n", myHero.Name, myHero.Power)
	fmt.Println("Инвентарь:", myHero.Inventory)
}
