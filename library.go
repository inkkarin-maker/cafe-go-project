package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Book struct {
	Name string `json:"book_name"`
	Page int    `json:"book_page"`
}
type Journal struct {
	Name string `json:"journal_name"`
	Size int    `json:"journal_size"`
}
type SpaceMap struct {
	Region string `json:"region_name"`
	Danger int    `json:"danger_level"`
}
type Readable interface {
	GetInfo() string
}

func (b Book) GetInfo() string {
	return fmt.Sprintf("Книга: %s, страниц: %d", b.Name, b.Page)
}
func (j Journal) GetInfo() string {
	return fmt.Sprintf("Журнал: %s, размер: %d", j.Name, j.Size)
}
func (s SpaceMap) GetInfo() string {
	return fmt.Sprintf("Карта сектора: %s, уровень опасности: %d/10", s.Region, s.Danger)
}
func handleLibrary(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	list := []Readable{
		Book{Name: "Interstellar", Page: 300},
		Journal{Name: "NASA", Size: 25},
		SpaceMap{Region: "Туманность Андромеды", Danger: 8},
	}
	fmt.Println("--- Содержимое библиотеки для пользователя ---")
	for _, item := range list {
		fmt.Println(item.GetInfo())
	}
	json.NewEncoder(w).Encode(list)
}
func main() {
	http.HandleFunc("/library", handleLibrary)
	fmt.Println("🚀 Сервер запущен на http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка сервера:", err)
	}
}
