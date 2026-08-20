package storage

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseMoney(input string) (int64, error) {
	s := strings.TrimSpace(input)
	if s == "" {
		return 0, nil
	}

	negative := false
	if strings.HasPrefix(s, "(") && strings.HasSuffix(s, ")") {
		negative = true
		s = strings.TrimPrefix(strings.TrimSuffix(s, ")"), "(")
	}

	var b strings.Builder
	for _, r := range s {
		switch {
		case r >= '0' && r <= '9':
			b.WriteRune(r)
		case r == '.':
			b.WriteRune(r)
		case r == '-':
			negative = true
		case r == ',' || r == ' ' || r == '\t':
			continue
		default:
			continue
		}
	}

	clean := b.String()
	if clean == "" || clean == "." {
		return 0, fmt.Errorf("invalid money value %q", input)
	}

	parts := strings.Split(clean, ".")
	if len(parts) > 2 {
		return 0, fmt.Errorf("invalid money value %q", input)
	}

	whole := parts[0]
	if whole == "" {
		whole = "0"
	}

	units, err := strconv.ParseInt(whole, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid money value %q: %w", input, err)
	}

	cents := int64(0)
	if len(parts) == 2 {
		frac := parts[1]
		if len(frac) > 2 {
			return 0, fmt.Errorf("money value %q has more than two decimal places", input)
		}
		for len(frac) < 2 {
			frac += "0"
		}
		cents, err = strconv.ParseInt(frac, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("invalid money value %q: %w", input, err)
		}
	}

	value := units*100 + cents
	if negative {
		value *= -1
	}
	return value, nil
}

func FormatMoney(cents int64) string {
	sign := ""
	if cents < 0 {
		sign = "-"
		cents *= -1
	}

	units := cents / 100
	remainder := cents % 100

	unitText := strconv.FormatInt(units, 10)
	for i := len(unitText) - 3; i > 0; i -= 3 {
		unitText = unitText[:i] + "," + unitText[i:]
	}

	return fmt.Sprintf("%s%s.%02d", sign, unitText, remainder)
}
