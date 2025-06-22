package validator

import (
	"fmt"
	"strings"
	"unicode"
)

func IsValidCNPJ(cnpj string) bool {
	cnpj = cleanCNPJ(cnpj)

	if len(cnpj) != 14 || isRepeated(cnpj) {
		return false
	}

	d1 := calculateCNPJDigit(cnpj[:12])
	d2 := calculateCNPJDigit(cnpj[:12] + d1)

	return cnpj[12:13] == d1 && cnpj[13:14] == d2
}

func cleanCNPJ(cnpj string) string {
	var b strings.Builder
	for _, r := range cnpj {
		if unicode.IsDigit(r) {
			b.WriteRune(r)
		}
	}
	return b.String()
}

func calculateCNPJDigit(base string) string {
	var weights = []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	if len(base) == 13 {
		weights = append([]int{6}, weights...)
	}

	sum := 0
	for i, r := range base {
		sum += int(r-'0') * weights[i]
	}
	d := 11 - (sum % 11)
	if d >= 10 {
		return "0"
	}
	return fmt.Sprintf("%d", d)
}
