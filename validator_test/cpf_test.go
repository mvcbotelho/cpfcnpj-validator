package validator_test

import (
	"testing"

	"github.com/mvcbotelho/cpfcnpj-validator/validator"
)

func TestIsValidCPF(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Valid CPF clean", "52998224725", true},
		{"Valid CPF formatted", "529.982.247-25", true},
		{"Invalid CPF digits", "52998224726", false},
		{"Invalid CPF repeated", "11111111111", false},
		{"Invalid CPF chars", "abc.def.ghi-jk", false},
		{"Invalid CPF short", "123", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := validator.IsValidCPF(tt.input)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
