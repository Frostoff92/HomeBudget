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

func TestParsePurchaseValid(t *testing.T) {
	record := []string{
		"TEST-001",
		"2026-11-28",
		"Lidl",
		"12,50",
		"groceries",
		"regular",
		"test.pdf",
		"Test purchase",
	}

	purchase, err := parsePurchase(record)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if purchase.AmountCents != 1250 {
		t.Errorf("Expected 1250, got %d", purchase.AmountCents)
	}

	if purchase.ID != "TEST-001" {
		t.Errorf("Expected ID TEST-001, got %s", purchase.ID)
	}

	if purchase.Store != record[2] {
		t.Errorf("Expected store %s, got %s", record[2], purchase.Store)
	}

	if purchase.Category != record[4] {
		t.Errorf("Expected category %s, got %s", record[4], purchase.Category)
	}

	if purchase.ExpenseType != ExpenseType(record[5]) {
		t.Errorf("Unknown expense type %s, got %s", record[5], purchase.ExpenseType)
	}

	if purchase.SourceFile != record[6] {
		t.Errorf("Not found source files %s, got %s", record[6], purchase.SourceFile)
	}

	if purchase.Notes != record[7] {
		t.Errorf("Notes not found, but its not required, got %s", purchase.Notes)
	}

	if purchase.Date.Format("2006-01-02") != record[1] {
		t.Errorf("Expected time %s, got %s", record[1], purchase.Date.Format("2006-01-02"))
	}
}

func TestParsePurchaseInvalidLength(t *testing.T) {
	record := []string{
		"TEST-001",
		"2026-11-28",
		"Lidl",
	}
	_, err := parsePurchase(record)

	if err == nil {
		t.Errorf("Expected error for invalid record length, got nil")
	}
}

func TestParsePurchaseInvalidMoney(t *testing.T) {
	record := []string{
		"TEST-001",
		"2026-11-28",
		"Lidl",
		"abc",
		"groceries",
		"regular",
		"test.pdf",
		"Test purchase",
	}
	_, err := parsePurchase(record)

	if err == nil {
		t.Errorf("Expected error for invalid money, got nil")
	}
}

func TestParsePurchaseInvalidDate(t *testing.T) {
	record := []string{
		"TEST-001",
		"2026-02-30",
		"Lidl",
		"12,50",
		"groceries",
		"regular",
		"test.pdf",
		"Test purchase",
	}
	_, err := parsePurchase(record)

	if err == nil {
		t.Errorf("Expected error for invalid date, got nil")
	}
}

func TestParsePurchaseInvalidAmount(t *testing.T) {
	record := []string{
		"TEST-001",
		"2026-01-02",
		"Lidl",
		"0,00",
		"groceries",
		"regular",
		"test.pdf",
		"Test purchase",
	}
	_, err := parsePurchase(record)

	if err == nil {
		t.Errorf("Expected error for amount must more than 0, got nil")
	}
}

func TestIsEmptyRecordEmpty(t *testing.T) {
	record := []string{
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
	}
	result := isEmptyRecord(record)
	if result != true {
		t.Errorf("Expected true, got %t", result)
	}
}

func TestIsEmptyRecordNonEmpty(t *testing.T) {
	record := []string{
		"TEST-001",
		"2026-11-28",
		"Lidl",
		"12,50",
		"groceries",
		"regular",
		"test.pdf",
		"Test purchase",
	}
	result := isEmptyRecord(record)
	if result != false {
		t.Errorf("Expected false, got %t", result)
	}
}

func TestIsEmptyRecordPartial(t *testing.T) {
	record := []string{
		"",
		"",
		"Lidl",
		"",
		"",
		"",
		"",
		"",
	}
	result := isEmptyRecord(record)
	if result != false {
		t.Errorf("Expected false, got %t", result)
	}
}
