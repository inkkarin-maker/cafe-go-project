package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

var count int

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	query := r.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Unknown"
	}
	count++
	os.WriteFile("count.txt", []byte(strconv.Itoa(count)), 0644)
	fmt.Fprintf(w, "<h1>Привет, %s!</h1>", name)
	fmt.Fprintf(w, "<p>Ты зашел на сервер в %d-й раз.</p>", count)
	fmt.Fprintf(w, "<h1>Главная страница</h1>")
	fmt.Fprintf(w, "<p style='color: red;'>Тут очень жарко!</p>")
	fmt.Fprintf(w, "<hr>")
	fmt.Fprintf(w, "<p>Вы — наш посетитель №%d!</p>", count)
}
func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Обо мне</h1><p>Я создаю роботов и теперь еще и сайты на Go!</p>")
}
func main() {
	data, err := os.ReadFile("count.txt")
	if err == nil {
		count, _ = strconv.Atoi(string(data))
	}
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	fmt.Println("Сервер запущен! \nГлавная: http://localhost:8080 \nО себе: http://localhost:8080/about")
	http.ListenAndServe(":8081", nil)
}
