package storage

import "testing"

func TestParseMoney(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int64
	}{
		{name: "plain cents", input: "12.34", want: 1234},
		{name: "negative", input: "-12.34", want: -1234},
		{name: "parentheses", input: "(12.34)", want: -1234},
		{name: "commas", input: "1,234.50", want: 123450},
		{name: "currency", input: "R 1,234.50", want: 123450},
		{name: "whole", input: "42", want: 4200},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseMoney(tt.input)
			if err != nil {
				t.Fatalf("ParseMoney() error = %v", err)
			}
			if got != tt.want {
				t.Fatalf("ParseMoney() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestParseMoneyRejectsTooManyDecimalPlaces(t *testing.T) {
	if _, err := ParseMoney("12.345"); err == nil {
		t.Fatal("ParseMoney() error = nil, want error")
	}
}

func TestFormatMoney(t *testing.T) {
	got := FormatMoney(-123456789)
	want := "-1,234,567.89"
	if got != want {
		t.Fatalf("FormatMoney() = %q, want %q", got, want)
	}
}
