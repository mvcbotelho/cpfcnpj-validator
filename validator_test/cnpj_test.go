package validator_test

import (
	"testing"

	"github.com/mvcbotelho/cpfcnpj-validator/validator"
)

func TestIsValidCNPJ(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Valid CNPJ clean", "11222333000181", true},
		{"Valid CNPJ formatted", "11.222.333/0001-81", true},
		{"Invalid CNPJ digits", "11222333000180", false},
		{"Invalid CNPJ repeated", "00000000000000", false},
		{"Invalid CNPJ chars", "abcd/efgh-ijkl", false},
		{"Invalid CNPJ short", "123", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := validator.IsValidCNPJ(tt.input)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
