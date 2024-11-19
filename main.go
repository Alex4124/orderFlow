package main

import (
	"encoding/json"
	"fmt"
)

type Customer struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Item struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type Order struct {
	OrderID  int      `json:"order_id"`
	Customer Customer `json:"customer"`
	Items    []Item   `json:"items"`
}

type Report struct {
	OrderID      int      `json:"order_id"`
	CustomerName string   `json:"customer_name"`
	TotalCost    float64  `json:"total_cost"`
	Items        []string `json:"items"`
}

// Функции для работы с JSON
func calculateTotalCost(items []Item) (float64, error) {
	var totalCost float64
	for _, item := range items {
		if item.Price <= 0 || item.Quantity <= 0 {
			return 0, fmt.Errorf("некорректная цена или количество у товара: %s", item.Name)
		}
		totalCost += item.Price * float64(item.Quantity)
	}
	return totalCost, nil
}

func extractItemNames(items []Item) []string {
	itemList := make([]string, len(items))
	for i, item := range items {
		itemList[i] = item.Name
	}
	return itemList
}

func generateReport(order Order) Report {
	if order.OrderID <= 0 {
		fmt.Println("Предупреждение: заказ имеет некорректный идентификатор.")
	}
	if len(order.Items) == 0 {
		fmt.Println("Предупреждение: заказ не содержит товаров.")
	}

	customerName := order.Customer.Name
	if customerName == "" {
		customerName = "Unknown Customer"
	}

	totalCost, err := calculateTotalCost(order.Items)
	if err != nil {
		fmt.Printf("Ошибка при расчете общей стоимости: %v\n", err)
		totalCost = 0
	}
	return Report{
		OrderID:      order.OrderID,
		CustomerName: customerName,
		TotalCost:    totalCost,
		Items:        extractItemNames(order.Items),
	}
}

func main() {
	// Пример JSON
	jsonData := `{
		"order_id": 12345,
		"customer": {
			"name": "John Doe",
			"email": "john.doe@example.com"
		},
		"items": [
			{"name": "Laptop", "price": 1000.50, "quantity": 1},
			{"name": "Mouse", "price": 25.00, "quantity": 2},
			{"name": "Keyboard", "price": 45.00, "quantity": 1}
		]
	}`

	var order Order

	// Десериализация JSON
	err := json.Unmarshal([]byte(jsonData), &order)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	// Генерация отчета
	report := generateReport(order)

	// Сериализация отчета в JSON
	reportJSON, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	// Вывод отчета
	fmt.Println(string(reportJSON))
}
