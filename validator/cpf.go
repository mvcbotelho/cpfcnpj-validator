package validator

func IsValidCPF(cpf string) bool {
	// Remove caracteres não numéricos
	cpf = cleanDocument(cpf)

	if len(cpf) != 11 || isRepeated(cpf) || isAllZeros(cpf) {
		return false
	}

	// Pesos para o primeiro dígito verificador
	weights1 := []int{10, 9, 8, 7, 6, 5, 4, 3, 2}
	d1 := calculateDigit(cpf[:9], weights1)

	// Pesos para o segundo dígito verificador
	weights2 := []int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2}
	d2 := calculateDigit(cpf[:9]+d1, weights2)

	return cpf[9:10] == d1 && cpf[10:11] == d2
}
