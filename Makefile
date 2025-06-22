# Nome das imagens
IMAGE_FINAL=cpfcnpj-validator
IMAGE_BUILDER=cpfcnpj-validator-builder

# Build da imagem final
build:
	docker build -t $(IMAGE_FINAL) .

# Build apenas da imagem builder (com Go instalado)
build-builder:
	docker build --target builder -t $(IMAGE_BUILDER) .

# Executar o binário para CPF
run-cpf:
	docker run --rm $(IMAGE_FINAL) --cpf 12345678909

# Executar o binário para CNPJ
run-cnpj:
	docker run --rm $(IMAGE_FINAL) --cnpj 12345678000195

# Rodar os testes com Go na imagem builder
test:
	docker build --target builder -t $(IMAGE_BUILDER) .
	docker run --rm -it --entrypoint go $(IMAGE_BUILDER) test ./...

# Limpar imagens intermediárias (opcional)
clean:
	docker rmi $(IMAGE_FINAL) $(IMAGE_BUILDER) || true
