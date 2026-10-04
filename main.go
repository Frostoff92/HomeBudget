package main

import (
	"fmt"
	"time"
)

type Purchace struct {
	ID          string
	Date        time.Time
	Store       string
	AmountCents int64
	Category    string
	ExpenseType string
}

func formatMoney(cents int64) string {
	return fmt.Sprintf("%d.%02d EUR", cents/100, cents%100)
}

func main() {
	purchace := Purchace{
		ID:          "HB-000001",
		Date:        time.Date(2026, 9, 19, 0, 0, 0, 0, time.Local),
		Store:       "Lidl",
		AmountCents: 1692,
		Category:    "groceries",
		ExpenseType: "regular",
	}

	fmt.Printf(
		"%s | %s | %s | %s | %s | %s\n",
		purchace.ID,
		purchace.Date.Format("2006-01-02"),
		purchace.Store,
		formatMoney(purchace.AmountCents),
		purchace.Category,
		purchace.ExpenseType,
	)
}
