package main

import "fmt"

// 1. Создаем интерфейс (стандарт)
type Payable interface {
	GetBalance() float64
}

// 2. Обычный счет
type BankAccount struct {
	Owner   string
	Balance float64
}

func (b BankAccount) GetBalance() float64 {
	return b.Balance
}

// 3. Крипто-кошелек
type CryptoWallet struct {
	CoinName string
	Amount   float64
	Price    float64 // цена за 1 монету
}

func (c CryptoWallet) GetBalance() float64 {
	return c.Amount * c.Price
}

// 4. Универсальная функция, которой плевать, что перед ней — банк или крипта
func printBalance(p Payable) {
	fmt.Printf("На счету сейчас: %.2f\n", p.GetBalance())
}

func main3() {
	myBank := BankAccount{Owner: "Inkar", Balance: 5000}
	myCrypto := CryptoWallet{CoinName: "Bitcoin", Amount: 0.5, Price: 40000}

	// Функция работает с обоими типами!
	printBalance(myBank)
	printBalance(myCrypto)
}
