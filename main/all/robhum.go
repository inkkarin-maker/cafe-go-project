package main

import (
	"fmt"
)

type Worker interface {
	Work() string
}
type Human struct {
	Name string
}
type Robot struct {
	ID int
}

func (r Robot) Work() string {
	return fmt.Sprintf("Робот #%d закончил сварку деталей.", r.ID)
}
func (h Human) Work() string {
	return fmt.Sprintf("Сотрудник %s заполнил отчеты.", h.Name)
}

func main() {
	logs := make(chan string)
	staff := make(map[string]Worker)
	staff["robot1"] = Robot{ID: 101}
	staff["human1"] = Human{Name: "Batyr"}
	for _, worker := range staff {
		go func(w Worker) {
			logs <- w.Work()
		}(worker)
	}
	fmt.Println("--- Сбор отчетов с фабрики ---")
	result1 := <-logs
	result2 := <-logs
	fmt.Println(result1)
	fmt.Println(result2)
}
