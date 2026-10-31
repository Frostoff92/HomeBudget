package main

import (
	"math"
	"testing"
	"time"
)

func TestParseMoney(t *testing.T) {
	got, err := parseMoney("38,2")

	if err != nil {
		t.Fatal(err)
	}

	if got != 3820 {
		t.Errorf("expected 3820, got %d", got)
	}
}

func TestParseMoneyInvalid(t *testing.T) {
	_, err := parseMoney("38,234")

	if err == nil {
		t.Error("expected parsing error, got nil")
	}
}

func TestParseMoneyTwoDecimals(t *testing.T) {
	got, err := parseMoney("10,44")

	if err != nil {
		t.Fatal(err)
	}

	if got != 1044 {
		t.Errorf("expected 1044, got %d", got)
	}
}

func TestParseMoneyNegative(t *testing.T) {
	_, err := parseMoney("-2,05")

	if err == nil {
		t.Errorf("expected error for negative amount, got nil")
	}
}

func TestParseMoneyNegativeZero(t *testing.T) {
	_, err := parseMoney("-0,05")

	if err == nil {
		t.Errorf("expected error for negative amount, got nil")
	}
}

func TestParseMoneyZero(t *testing.T) {
	got, err := parseMoney("0,00")

	if err != nil {
		t.Fatal(err)
	}

	if got != 0 {
		t.Errorf("expected 0, got %d", got)
	}
}

func TestPurchaseZeroAmount(t *testing.T) {
	purchase := Purchase{
		ID:          "HB-TEST-001",
		Date:        time.Date(2026, 10, 31, 0, 0, 0, 0, time.UTC),
		Store:       "Lidl",
		AmountCents: 0,
		Category:    "groceries",
		ExpenseType: ExpenseRegular,
	}

	err := purchase.Validate()
	if err == nil {
		t.Error("expected validation error, err nil")
	}
}

func TestPurchaseValid(t *testing.T) {
	purchase := Purchase{
		ID:          "HB-TEST-001",
		Date:        time.Date(2026, 10, 31, 0, 0, 0, 0, time.UTC),
		Store:       "Lidl",
		AmountCents: 205,
		Category:    "groceries",
		ExpenseType: ExpenseRegular,
	}

	err := purchase.Validate()
	if err != nil {
		t.Errorf("expected valid purchase, got error: %v", err)
	}
}

func TestParseMoneyOverflow(t *testing.T) {
	_, err := parseMoney("92233720368547758,08")

	if err == nil {
		t.Error("expected overflow error, got nil")
	}
}

func TestParseMoneyMaxInt64(t *testing.T) {
	got, err := parseMoney("92233720368547758,07")

	if err != nil {
		t.Fatal(err)
	}

	if got != math.MaxInt64 {
		t.Errorf("expected %d, got %d", math.MaxInt64, got)
	}
}

func TestParseMoneyNegativeCents(t *testing.T) {
	_, err := parseMoney("12,-5")

	if err == nil {
		t.Error("expected error for negative amount, nil")
	}
}

func TestParseMoneyPlusCents(t *testing.T) {
	_, err := parseMoney("12,+5")

	if err == nil {
		t.Error("expected error for plus cents, nil")
	}
}
