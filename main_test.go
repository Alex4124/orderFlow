package main

import (
	"testing"
)

func TestCalculateTotalCost(t *testing.T) {
	tests := []struct {
		name    string
		items   []Item
		want    float64
		wantErr bool
	}{
		{
			name: "Valid items",
			items: []Item{
				{"Laptop", 1000.50, 1},
				{"Mouse", 25.00, 2},
			},
			want:    1050.50,
			wantErr: false,
		},
		{
			name: "Negative price",
			items: []Item{
				{"Laptop", -1000.50, 1},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Zero quantity",
			items: []Item{
				{"Mouse", 25.00, 0},
			},
			want:    0,
			wantErr: true,
		},
		{
			name:    "Empty item list",
			items:   []Item{},
			want:    0,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calculateTotalCost(tt.items)
			if (err != nil) != tt.wantErr {
				t.Errorf("calculateTotalCost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("calculateTotalCost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtractItemNames(t *testing.T) {
	tests := []struct {
		name  string
		items []Item
		want  []string
	}{
		{
			name: "Multiple items",
			items: []Item{
				{"Laptop", 1000.50, 1},
				{"Mouse", 25.00, 2},
			},
			want: []string{"Laptop", "Mouse"},
		},
		{
			name:  "Empty list",
			items: []Item{},
			want:  []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := extractItemNames(tt.items)
			if len(got) != len(tt.want) {
				t.Errorf("extractItemNames() = %v, want %v", got, tt.want)
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("extractItemNames() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestGenerateReport(t *testing.T) {
	tests := []struct {
		name  string
		order Order
		want  Report
	}{
		{
			name: "Valid order",
			order: Order{
				OrderID: 12345,
				Customer: Customer{
					Name: "John Doe",
				},
				Items: []Item{
					{"Laptop", 1000.50, 1},
					{"Mouse", 25.00, 2},
				},
			},
			want: Report{
				OrderID:      12345,
				CustomerName: "John Doe",
				TotalCost:    1050.50,
				Items:        []string{"Laptop", "Mouse"},
			},
		},
		{
			name: "Empty customer name",
			order: Order{
				OrderID: 12346,
				Customer: Customer{
					Name: "",
				},
				Items: []Item{
					{"Keyboard", 45.00, 1},
				},
			},
			want: Report{
				OrderID:      12346,
				CustomerName: "Unknown Customer",
				TotalCost:    45.00,
				Items:        []string{"Keyboard"},
			},
		},
		{
			name: "Empty items list",
			order: Order{
				OrderID: 12347,
				Customer: Customer{
					Name: "Jane Doe",
				},
				Items: []Item{},
			},
			want: Report{
				OrderID:      12347,
				CustomerName: "Jane Doe",
				TotalCost:    0,
				Items:        []string{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateReport(tt.order)
			if got.OrderID != tt.want.OrderID || got.CustomerName != tt.want.CustomerName || got.TotalCost != tt.want.TotalCost {
				t.Errorf("generateReport() = %v, want %v", got, tt.want)
			}
			if len(got.Items) != len(tt.want.Items) {
				t.Errorf("generateReport() items = %v, want %v", got.Items, tt.want.Items)
				return
			}
			for i := range got.Items {
				if got.Items[i] != tt.want.Items[i] {
					t.Errorf("generateReport() items[%d] = %v, want %v", i, got.Items[i], tt.want.Items[i])
				}
			}
		})
	}
}
