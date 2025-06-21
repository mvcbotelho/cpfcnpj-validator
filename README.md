# 🧾 Validador de CPF/CNPJ em Go

Este projeto é uma ferramenta simples e eficiente para validação de CPF e CNPJ, escrita em Go e totalmente containerizada com Docker. Ideal para aprendizado, reuso como biblioteca ou uso em linha de comando.

---

## 🚀 Como executar com Docker

### 1. Build da imagem:

```bash
docker build -t cpfcnpj-validator .
```

### 2. Rodar com CPF:

```bash
docker run --rm cpfcnpj-validator --cpf 12345678909
```

### 3. Rodar com CNPJ:

```bash
docker run --rm cpfcnpj-validator --cnpj 12345678000195
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
├── Dockerfile
├── go.mod
└── README.md
```

---

## ✅ O que já funciona

- Interface de linha de comando para validação de CPF e CNPJ
- Dockerfile multi-stage leve
- Estrutura modular e organizada para evoluir facilmente

---

## 💡 Possíveis melhorias

- Validação completa com dígitos verificadores
- Testes unitários com `go test`
- Exportar como biblioteca Go para reuso em outros projetos
- API REST para validação via HTTP
- Upload e validação em lote (CSV)

---

## 🤝 Contribuindo

Sinta-se livre para abrir issues, sugerir melhorias ou enviar PRs!

---

## 🧠 Autor

**Marcus Botelho**  

---

## 📜 Licença

MIT License. Veja o arquivo `LICENSE` para mais detalhes.
