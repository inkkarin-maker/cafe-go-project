package main

import (
	"fmt"
	"sync"
	"time"
)

type File struct {
	Name string
	Size int
}

func scanFiles(name string, files chan File, wg *sync.WaitGroup) {
	defer wg.Done()
	for f := range files {
		fmt.Println(name, "проверяет файл", f.Name)
		time.Sleep(time.Duration(f.Size) * time.Millisecond * 100)
		fmt.Println(f.Name, "проверен!")
	}
}

func main() {
	var wg sync.WaitGroup
	files := make(chan File)
	wg.Add(2)
	go scanFiles("Сканер А", files, &wg)
	go scanFiles("Сканер B", files, &wg)
	files <- File{Name: "Report", Size: 25}
	files <- File{Name: "CV", Size: 78}
	files <- File{Name: "Bank", Size: 125}
	files <- File{Name: "Names", Size: 15}
	close(files)
	wg.Wait()
}
