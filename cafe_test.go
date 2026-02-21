package main

import "testing"

func TestApplyDiscount(t *testing.T) {
	order := Order{
		Name:  "TestUser",
		Drink: "Latte",
		Price: 1000,
	}
	discountPercent := 10
	order.ApplyDiscount(discountPercent)
	expectedPrice := 900
	if order.Price != expectedPrice {
		t.Errorf("Ошибка в ApplyDiscount: ожидалось %d, но получили %d", expectedPrice, order.Price)
	}
}
