# 🧾 Validador de CPF/CNPJ em Go

Este projeto é uma ferramenta simples e eficiente para validação de CPF e CNPJ, escrita em Go e totalmente containerizada com Docker. Ideal para aprendizado, reuso como biblioteca ou uso em linha de comando.

---

## 🚀 Como executar

### Com Docker

1. **Build da imagem:**

```bash
make docker-build
```

2. **Rodar validação de CPF:**

```bash
make docker-run-cpf
```

3. **Rodar validação de CNPJ:**

```bash
make docker-run-cnpj
```

4. **Rodar testes na imagem Docker:**

```bash
make docker-test
```

### Localmente (sem Docker)

1. **Rodar validação de CPF:**

```bash
make run-local
```

2. **Rodar testes:**

```bash
make test-local
```

3. **Lint do código (requer golangci-lint):**

```bash
make lint
```

4. **Formatar código:**

```bash
make format
```

---

## 📁 Estrutura do projeto

```
cpfcnpj-validator/
├── main.go
├── validator/
│   ├── cpf.go
│   ├── cnpj.go
│   └── utils.go
├── validator_test/
│   ├── cpf_test.go
│   └── cnpj_test.go
├── Dockerfile
├── Makefile
├── go.mod
└── README.md
```

---

## ✅ O que já funciona

- Interface de linha de comando para validação de CPF e CNPJ
- Dockerfile multi-stage seguro, rodando como usuário não-root e com HEALTHCHECK
- Estrutura modular e organizada para evoluir facilmente
- Funções utilitárias para limpeza, repetição e cálculo de dígitos
- Validação real com dígitos verificadores
- Testes unitários com `go test`
- Exportar como biblioteca Go para reuso em outros projetos
- Makefile completo com targets para build, run, test, lint e format

---

## 💡 Possíveis melhorias

- API REST para validação via HTTP
- Upload e validação em lote (CSV)

---

## 📦 Como usar como biblioteca

Você pode importar diretamente no seu projeto Go:

```go
import "github.com/mvcbotelho/cpfcnpj-validator/validator"

func main() {
    if validator.IsValidCPF("12345678909") {
        fmt.Println("CPF válido")
    }
}
```

### Testar localmente:

Adicione no `go.mod` do seu projeto:

```
replace github.com/mvcbotelho/cpfcnpj-validator => ../cpfcnpj-validator
```

---

## 🆕 Melhorias recentes

- Refatoração: código duplicado removido, funções utilitárias criadas em `utils.go`
- Dockerfile mais seguro: multi-stage, usuário não-root, HEALTHCHECK
- Makefile expandido: targets para build, run, test, lint, format, local e Docker
- Testes organizados em `validator_test/`
- Validação aprimorada para CPF e CNPJ

---

## 🤝 Contribuindo

Sinta-se livre para abrir issues, sugerir melhorias ou enviar PRs!

---

## 🧠 Autor

**Marcus Botelho** 

- GitHub: [@mvcbotelho](https://github.com/mvcbotelho)
- LinkedIn: [@mvcbotelho](https://linkedin.com/in/mvcbotelho)

---

## 📜 Licença

MIT License. Veja o arquivo `LICENSE` para mais detalhes.
