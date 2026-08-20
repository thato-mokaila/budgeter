package storage

import (
	"fmt"
	"strings"
	"time"
)

const DateLayout = "2006-01-02"

var dateLayouts = []string{
	DateLayout,
	"2006/01/02",
	"02/01/2006",
	"02-01-2006",
	time.RFC3339,
}

func Today() string {
	return time.Now().Format(DateLayout)
}

func NormalizeDate(input string) (string, error) {
	s := strings.TrimSpace(input)
	if s == "" {
		return "", nil
	}

	for _, layout := range dateLayouts {
		t, err := time.Parse(layout, s)
		if err == nil {
			return t.Format(DateLayout), nil
		}
	}

	return "", fmt.Errorf("invalid date %q, expected YYYY-MM-DD", input)
}

func CurrentMonthRange() (string, string) {
	start, end, _, _ := MonthRange(Today())
	return start, end
}

func MonthRange(input string) (string, string, string, error) {
	date, err := NormalizeDate(input)
	if err != nil {
		return "", "", "", err
	}
	if date == "" {
		date = Today()
	}

	t, err := time.Parse(DateLayout, date)
	if err != nil {
		return "", "", "", fmt.Errorf("invalid date %q, expected YYYY-MM-DD", input)
	}

	start := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
	end := start.AddDate(0, 1, 0)
	return start.Format(DateLayout), end.Format(DateLayout), start.Format("2006-01"), nil
}
