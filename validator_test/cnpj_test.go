package validator_test

import (
	"testing"

	"github.com/mvcbotelho/cpfcnpj-validator/validator"
)

func TestValidCNPJ(t *testing.T) {
	valid := []string{
		"45.723.174/0001-10",
		"04.252.011/0001-10",
		"11222333000181",
	}

	for _, cnpj := range valid {
		if !validator.IsValidCNPJ(cnpj) {
			t.Errorf("CNPJ deveria ser válido: %s", cnpj)
		}
	}
}

func TestInvalidCNPJ(t *testing.T) {
	invalid := []string{
		"45.723.174/0001-00",
		"00.000.000/0000-00",
		"11111111111111",
		"abcdefghijklll",
		"11222333000180", // dígito verificador errado
	}

	for _, cnpj := range invalid {
		if validator.IsValidCNPJ(cnpj) {
			t.Errorf("CNPJ deveria ser inválido: %s", cnpj)
		}
	}
}
