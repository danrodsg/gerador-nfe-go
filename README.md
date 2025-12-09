# 🧾 Gerador de Notas Fiscais Fake (GoLang)

[![Go](https://github.com/golang/go/blob/master/assets/badge.svg)](https://golang.org/)
[![JSON](https://img.shields.io/badge/Data%20Format-JSON-lightgray.svg)](https://www.json.org/json-en.html)

Este repositório contém uma API simples em **GoLang** projetada para **gerar dados de Notas Fiscais Eletrônicas (NFe) simuladas** em formato JSON. O objetivo é fornecer dados estruturados para **testes, desenvolvimento de front-ends** ou **simulações de integrações** sem depender de sistemas fiscais reais.

O projeto demonstra a modelagem de dados complexos usando `structs` e `json` tags em Go, e utiliza o pacote `net/http` da biblioteca padrão para servir o endpoint.

---

## 🛠️ Estrutura de Dados e Modelagem

O projeto utiliza structs bem definidas (vistas em `models/models.go`) para modelar o documento fiscal e suas partes, garantindo a correta serialização para JSON.

### 1. Modelo Principal (`NotaFiscal`)

A estrutura principal agrega as informações e usa *tags* JSON:

```go
type NotaFiscal struct {
    ChaveAcesso string       `json:"chave_acesso"`
    DataEmissao time.Time    `json:"data_emissao"`
    Emitente    Emitente     `json:"emitente"`
    Destinatario Destinatario `json:"destinatario"`
    Itens       []Item       `json:"itens"`
    Totais      Totais       `json:"totais"`
}

````
### 2. Campos de Destaque

As sub-estruturas definem os campos essenciais do documento, que são preenchidos por lógica no generator/generator.go:

- Estrutura	/Campo de Destaque	/Tipo	/Descrição
- 
- Emitente/	CNPJ	/string	/CNPJ do emissor da nota (simulado).
- 
- Destinatario/	CPF_CNPJ/	string	/CPF ou CNPJ do cliente (simulado).
- 
- Item	/ValorUnitario	/float64	P/reço unitário do produto.
- 
- Totais	/ValorTotalNF	/float64	/Valor total da nota (calculado a partir dos itens gerados).



## ✨ Funcionalidades (Endpoint)

A API expõe um único endpoint para a geração da nota fiscal.

- Método HTTP	/Endpoint	/Descrição
- GET/	/api/v1/generate	/Principal Endpoint. Executa a lógica de geração de dados e retorna um objeto JSON contendo uma Nota Fiscal simulada.

  ##⚙️ Como Executar o Projeto

 ### 1. Pré-requisitos
Golang: Versão 1.16 ou superior.

Git: Para clonar o repositório.

### 2. Clonar o Repositório

```
bash
git clone github.com/danrodsg/gerador-nfe-go.git
cd gerador-nfe-go

````
### 3. Executar a API
````
bash
 O ponto de entrada principal é o cmd/api/main.go
go run cmd/api/main.go

````
O servidor estará rodando em http://localhost:8080.

## 🧪 Teste Rápido

Para obter uma Nota Fiscal simulada, use o curl:
````
bash
curl -X GET http://localhost:8080/api/v1/generate
