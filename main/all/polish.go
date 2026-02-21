package main

import (
	"fmt"
	"sync"
	"time"
)

type Detail struct {
	Part string
	Time int
}

func polish(name string, polishers chan Detail, wg *sync.WaitGroup) {
	defer wg.Done()
	for p := range polishers {
		fmt.Println(name, "начал покраску ", p.Part)
		time.Sleep(time.Duration(p.Time) * time.Millisecond * 10)
		fmt.Println("Покраска ", p.Part, "окончена!")
	}
}
func main() {
	var wg sync.WaitGroup
	details := make(chan Detail, 10)
	wg.Add(3)
	go polish("Робот №1", details, &wg)
	go polish("Робот №2", details, &wg)
	go polish("Робот №3", details, &wg)
	details <- Detail{Part: "Hand", Time: 45}
	details <- Detail{Part: "Foot", Time: 15}
	details <- Detail{Part: "Body", Time: 120}
	details <- Detail{Part: "Head", Time: 36}
	close(details)
	wg.Wait()
}
