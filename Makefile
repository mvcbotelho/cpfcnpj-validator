build:
	docker build -t cpfcnpj-validator .

run-cpf:
	docker run --rm cpfcnpj-validator --cpf 12345678909

run-cnpj:
	docker run --rm cpfcnpj-validator --cnpj 12345678000195
