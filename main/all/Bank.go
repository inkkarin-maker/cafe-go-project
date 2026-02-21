package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

// 1. Структура
type BankAccount1 struct {
	ID      int     `json:"id"`
	Owner   string  `json:"owner"`
	Balance float64 `json:"balance"`
	PinCode int     `json:"pin_code"` // Добавили тег для регистрации
}

var db *sql.DB

// --- ТУТ ТВОИ ФУНКЦИИ-ОБРАБОТЧИКИ ---

func getBalance(w http.ResponseWriter, r *http.Request) { /* ... код ... */ }

func withdrawMoney(w http.ResponseWriter, r *http.Request) { /* ... код ... */ }

// ВСТАВЛЯЙ СЮДА:
func createAccount(w http.ResponseWriter, r *http.Request) {
	var newAcc BankAccount
	err := json.NewDecoder(r.Body).Decode(&newAcc)
	if err != nil {
		http.Error(w, "Ошибка декодирования JSON", 400)
		return
	}

	query := `INSERT INTO accounts (owner, balance, pin_code) 
              VALUES ($1, $2, $3) RETURNING id`

	err = db.QueryRow(query, newAcc.Owner, newAcc.Balance, newAcc.PinCode).Scan(&newAcc.ID)

	if err != nil {
		log.Println("Ошибка БД:", err)
		http.Error(w, "Ошибка при создании в БД", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newAcc)
}

// --- ГЛАВНАЯ ФУНКЦИЯ ---

func main2() {
	// ... подключение к БД ...

	// ОБЯЗАТЕЛЬНО ДОБАВЬ ЭТУ СТРОЧКУ ЗДЕСЬ:
	http.HandleFunc("/balance", getBalance)
	http.HandleFunc("/withdraw", withdrawMoney)
	http.HandleFunc("/create", createAccount) // <--- Регистрация нового маршрута

	log.Fatal(http.ListenAndServe(":8080", nil))
}
