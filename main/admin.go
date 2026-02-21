package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
)

type UserData struct {
	Name       string
	VisitCount int
	IsAdmin    bool
}

var count int

func Home(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("user_name")
	count++
	os.WriteFile("count.txt", []byte(strconv.Itoa(count)), 0644)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	isAdmin := (name == "admin")
	user := UserData{
		Name:       name,
		VisitCount: count,
		IsAdmin:    isAdmin,
	}
	tmpl, _ := template.ParseFiles("index.html")
	tmpl.Execute(w, user)
}
func main() {
	data, err := os.ReadFile("count.txt")
	if err == nil {
		count, _ = strconv.Atoi(string(data))
	}

	http.HandleFunc("/", Home)
	fmt.Println("Сервер запущен на http://localhost:8081")
	http.ListenAndServe(":8081", nil)
}
