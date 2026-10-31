package main

import (
	"fmt"
	"time"
)

type Purchase struct {
	ID          string
	Date        time.Time
	Store       string
	AmountCents int64
	Category    string
	ExpenseType ExpenseType
}

type ExpenseType string

const (
	ExpenseRegular ExpenseType = "regular"
	ExpenseOneTime ExpenseType = "one_time"
	ExpenseUnknown ExpenseType = "unknown"
)

func (e ExpenseType) IsValid() bool {
	switch e {
	case ExpenseRegular, ExpenseOneTime, ExpenseUnknown:
		return true
	default:
		return false
	}
}

func (p Purchase) Validate() error {
	if p.ID == "" {
		return fmt.Errorf("purchases ID is empty")
	}

	if p.Date.IsZero() {
		return fmt.Errorf("purchases date is empty")
	}

	if p.Store == "" {
		return fmt.Errorf("store is empty")
	}

	if p.AmountCents <= 0 {
		return fmt.Errorf("amount must be greater than zero")
	}

	if p.Category == "" {
		return fmt.Errorf("category is empty")
	}

	if !p.ExpenseType.IsValid() {
		return fmt.Errorf("invalid expense type: %s", p.ExpenseType)
	}

	return nil
}

func formatMoney(cents int64) string {
	return fmt.Sprintf("%d.%02d EUR", cents/100, cents%100)
}

func calculateTotal(purchases []Purchase) int64 {
	var total int64

	for _, purchase := range purchases {
		err := purchase.Validate()
		if err != nil {
			continue
		}

		total += purchase.AmountCents
	}

	return total
}

func calculateTotalStore(purchases []Purchase, store string) int64 {
	var totalStore int64

	for _, purchase := range purchases {
		err := purchase.Validate()
		if err != nil {
			continue
		}
		if store != purchase.Store {
			continue
		}

		totalStore += purchase.AmountCents
	}

	return totalStore
}

func main() {
	purchases := []Purchase{
		{
			ID:          "HB-000001",
			Date:        time.Date(2026, 9, 19, 0, 0, 0, 0, time.Local),
			Store:       "Lidl",
			AmountCents: 1692,
			Category:    "groceries",
			ExpenseType: ExpenseRegular,
		},
		{
			ID:          "HB-000002",
			Date:        time.Date(2026, 9, 13, 0, 0, 0, 0, time.Local),
			Store:       "Lidl",
			AmountCents: 0,
			Category:    "groceries",
			ExpenseType: ExpenseRegular,
		},
		{
			ID:          "HB-000003",
			Date:        time.Date(2026, 9, 25, 0, 0, 0, 0, time.Local),
			Store:       "Kaufland",
			AmountCents: 3064,
			Category:    "groceries",
			ExpenseType: ExpenseRegular,
		},
	}

	fmt.Println(len(purchases))

	fmt.Println("Total:", formatMoney(calculateTotal(purchases)))

	fmt.Println("Lidl total:", formatMoney(calculateTotalStore(purchases, "Lidl")))

	for _, purchase := range purchases {

		err := purchase.Validate()

		if err != nil {
			fmt.Println("Validation error:", err)
			continue
		}

		fmt.Println(purchase.ID)
	}

}
