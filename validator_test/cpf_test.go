package validator_test

import (
	"testing"

	"github.com/mvcbotelho/cpfcnpj-validator/validator"
)

func TestValidCPF(t *testing.T) {
	valid := []string{
		"123.456.789-09",
		"529.982.247-25",
		"11144477735",
	}

	for _, cpf := range valid {
		if !validator.IsValidCPF(cpf) {
			t.Errorf("CPF deveria ser válido: %s", cpf)
		}
	}
}

func TestInvalidCPF(t *testing.T) {
	invalid := []string{
		"123.456.789-00",
		"000.000.000-00",
		"11111111111",
		"99999999999",
		"abc.def.ghi-jk",
		"529.982.247-26",
	}

	for _, cpf := range invalid {
		if validator.IsValidCPF(cpf) {
			t.Errorf("CPF deveria ser inválido: %s", cpf)
		}
	}
}
