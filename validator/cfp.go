package validator

import (
	"strings"
	"unicode"
)

func IsValidCPF(cpf string) bool {
	// Remove caracteres não numéricos
	cpf = cleanCPF(cpf)

	if len(cpf) != 11 || isRepeated(cpf) {
		return false
	}

	d1 := calculateDigit(cpf[:9], 10)
	d2 := calculateDigit(cpf[:9]+d1, 11)

	return cpf[9:10] == d1 && cpf[10:11] == d2
}

func cleanCPF(cpf string) string {
	var b strings.Builder
	for _, r := range cpf {
		if unicode.IsDigit(r) {
			b.WriteRune(r)
		}
	}
	return b.String()
}

func isRepeated(s string) bool {
	for i := 1; i < len(s); i++ {
		if s[i] != s[0] {
			return false
		}
	}
	return true
}

func calculateDigit(base string, weight int) string {
	sum := 0
	for _, r := range base {
		sum += int(r-'0') * weight
		weight--
	}
	d := 11 - (sum % 11)
	if d >= 10 {
		return "0"
	}
	return string('0' + d)
}
