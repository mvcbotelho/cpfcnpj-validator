package validator

func IsValidCNPJ(cnpj string) bool {
	cnpj = cleanDocument(cnpj)

	if len(cnpj) != 14 || isRepeated(cnpj) || isAllZeros(cnpj) {
		return false
	}

	d1 := calculateCNPJDigit(cnpj[:12])
	d2 := calculateCNPJDigit(cnpj[:12] + d1)

	return cnpj[12:13] == d1 && cnpj[13:14] == d2
}

func calculateCNPJDigit(base string) string {
	var weights = []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	if len(base) == 13 {
		weights = append([]int{6}, weights...)
	}

	return calculateDigit(base, weights)
}
