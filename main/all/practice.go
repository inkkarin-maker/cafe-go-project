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
	name := r.URL.Query().Get("user_name")
	count++
	os.WriteFile("count.txt", []byte(strconv.Itoa(count)), 0644)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h2>Наш уютный сервер</h2>")

	if name == "" {
		fmt.Fprintf(w, `
			<p>Пожалуйста, представьтесь:</p>
			<form action="/" method="GET">
				<input type="text" name="user_name" placeholder="Ваше имя" required>
				<button type="submit">Войти на сервер</button>
			</form>
		`)
	} else {
		if name == "admin" {
			fmt.Fprintf(w, "<h1 style='color: blue;'>Добро пожаловать, Офицер!</h1>")
		} else {
			fmt.Fprintf(w, "<h1>Привет, %s!</h1>", name)
		}
		fmt.Fprintf(w, "<a href='/'>Назад к вводу имени</a>")
	}

	fmt.Fprintf(w, "<hr><p>Посетителей всего: %d</p>", count)
}
func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Обо мне</h1><p>Я создаю роботов и сайты на Go!</p>")
}
func main() {
	data, err := os.ReadFile("count.txt")
	if err == nil {
		count, _ = strconv.Atoi(string(data))
	}
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	fmt.Println("Сервер запущен! \nГлавная: http://localhost:8081 \nО себе: http://localhost:8081/about")
	http.ListenAndServe(":8081", nil)
}
