package main

import (
	"fmt"
	"sync"
	"time"
)

type Package struct {
	Destination string
	Weight      int
}

func deliever(name string, delievers chan Package, wg *sync.WaitGroup) {
	defer wg.Done()
	for d := range delievers {
		fmt.Println(name, "везет посылку в ", d.Destination)
		time.Sleep(time.Duration(d.Weight) * time.Millisecond * 1)
		fmt.Println("Посылка в ", d.Destination, "доставлена!")
	}
}
func main() {
	var wg sync.WaitGroup
	packages := make(chan Package, 10)
	wg.Add(3)
	go deliever("Курьeр №1", packages, &wg)
	go deliever("Курьeр №2", packages, &wg)
	go deliever("Курьeр №3", packages, &wg)
	packages <- Package{Destination: "Almaty", Weight: 250}
	packages <- Package{Destination: "Astana", Weight: 560}
	packages <- Package{Destination: "Atyrau", Weight: 850}
	packages <- Package{Destination: "Aktau", Weight: 960}
	close(packages)
	wg.Wait()
}
