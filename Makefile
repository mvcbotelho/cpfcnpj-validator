# Nome das imagens
IMAGE_FINAL=cpfcnpj-validator
IMAGE_BUILDER=cpfcnpj-validator-builder

# Build da imagem final
docker-build:
	docker build -t $(IMAGE_FINAL) .

# Build apenas da imagem builder (com Go instalado)
docker-build-builder:
	docker build --target builder -t $(IMAGE_BUILDER) .

# Executar o binário para CPF via Docker
docker-run-cpf:
	docker run --rm $(IMAGE_FINAL) --cpf 12345678909

# Executar o binário para CNPJ via Docker
docker-run-cnpj:
	docker run --rm $(IMAGE_FINAL) --cnpj 12345678000195

# Rodar os testes com Go na imagem builder
docker-test:
	docker build --target builder -t $(IMAGE_BUILDER) .
	docker run --rm -it --entrypoint go $(IMAGE_BUILDER) test ./...

# Limpar imagens intermediárias (opcional)
docker-clean:
	docker rmi $(IMAGE_FINAL) $(IMAGE_BUILDER) || true

# Rodar localmente (sem Docker)
run-local:
	go run main.go --cpf 12345678909

# Testar localmente (sem Docker)
test-local:
	go test ./...

# Lint (requer golangci-lint instalado)
lint:
	golangci-lint run ./...

# Format (padronizar código)
format:
	gofmt -s -w .
