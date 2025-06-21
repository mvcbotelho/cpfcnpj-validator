package main

import (
	"fmt"
	"os"

	"github.com/mvcbotelho/cpfcnpj-validator/validator"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Uso: validator --cpf 12345678909 ou --cnpj 12345678000195")
		return
	}

	arg := os.Args[1]
	value := os.Args[2]

	switch arg {
	case "--cpf":
		if validator.IsValidCPF(value) {
			fmt.Println("CPF válido")
		} else {
			fmt.Println("CPF inválido")
		}
	case "--cnpj":
		if validator.IsValidCNPJ(value) {
			fmt.Println("CNPJ válido")
		} else {
			fmt.Println("CNPJ inválido")
		}
	default:
		fmt.Println("Parâmetro inválido. Use --cpf ou --cnpj")
	}
}
