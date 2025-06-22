package validator

import (
	"fmt"
	"strings"
	"unicode"
)

// cleanDocument remove caracteres não numéricos de um documento
func cleanDocument(doc string) string {
	var b strings.Builder
	for _, r := range doc {
		if unicode.IsDigit(r) {
			b.WriteRune(r)
		}
	}
	return b.String()
}

// isRepeated verifica se todos os dígitos são iguais
func isRepeated(s string) bool {
	if len(s) == 0 {
		return false
	}
	for i := 1; i < len(s); i++ {
		if s[i] != s[0] {
			return false
		}
	}
	return true
}

// isAllZeros verifica se todos os dígitos são zero
func isAllZeros(s string) bool {
	for _, r := range s {
		if r != '0' {
			return false
		}
	}
	return true
}

// calculateDigit calcula um dígito verificador baseado em pesos
func calculateDigit(base string, weights []int) string {
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
