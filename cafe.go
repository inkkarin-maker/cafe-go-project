package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Order struct {
	Name  string `json:"customer_name"`
	Drink string `json:"drink_ordered"`
	Price int    `json:"final_price"`
}

type Item interface {
	GetPrice() int
}

func (o Order) GetPrice() int {
	return o.Price
}

type Pastry struct {
	Name  string
	Price int
}

func (p Pastry) GetPrice() int {
	return p.Price
}
func PrintTotal(goods []Item) {
	sum := 0
	for _, g := range goods {
		sum += g.GetPrice()
	}
	fmt.Printf("--- ИТОГО К ОПЛАТЕ: %d ---\n", sum)
}
func (o Order) PrintStatus() {
	fmt.Printf(">> [ЧЕК] Клиент: %s | Напиток: %s | Цена: %d\n", o.Name, o.Drink, o.Price)
}

func (o *Order) ApplyDiscount(percent int) {
	o.Price = o.Price - (o.Price * percent / 100)
}

func getPrice(menu map[string]int, drink string) (int, error) {
	price, ok := menu[drink]
	if !ok {
		return 0, fmt.Errorf("напиток %s не найден", drink)
	}
	return price, nil
}
func handleOrder(w http.ResponseWriter, r *http.Request) {
	menu := map[string]int{
		"Latte":    1600,
		"Espresso": 1200,
		"Tea":      1450,
	}
	drinkName := r.URL.Query().Get("drink")
	customerName := r.URL.Query().Get("name")
	if customerName == "" {
		customerName = "Guest"
	}
	price, ok := menu[drinkName]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Извините, напитка '%s' нет в меню.", drinkName)
		return
	}
	ord := Order{Name: customerName, Drink: drinkName, Price: price}
	ord.ApplyDiscount(10)
	w.Header().Set("Content-Type", "application/json")
	jsonResponse, err := json.Marshal(ord)
	if err != nil {
		http.Error(w, "Ошибка при создании JSON", http.StatusInternalServerError)
		return
	}
	w.Write(jsonResponse)
}

func main() {
	http.HandleFunc("/order", handleOrder)
	fmt.Println("🚀 Сервер запущен на http://localhost:8080")
	fmt.Println("Открой браузер и введи: http://localhost:8080/order?drink=Latte&name=Assem")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка сервера:", err)
	}
}
