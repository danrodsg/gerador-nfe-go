// internal/models/nfe.go

package models

import "time"

type NotaFiscal struct {
	ChaveAcesso  string       `json:"chave_acesso"`
	DataEmissao  time.Time    `json:"data_emissao"`
	Emitente     Emitente     `json:"emitente"`
	Destinatario Destinatario `json:"destinatario"`
	Itens        []Item       `json:"itens"`
	Totais       Totais       `json:"totais"`
}

type Emitente struct {
	CNPJ string `json:"cnpj"`
	Nome string `json:"nome"`
}

type Destinatario struct {
	CPF_CNPJ string `json:"cpf_cnpj"`
	Nome     string `json:"nome"`
}

type Item struct {
	Codigo        string  `json:"codigo"`
	Descricao     string  `json:"descricao"`
	Quantidade    int     `json:"quantidade"`
	ValorUnitario float64 `json:"valor_unitario"`
}

type Totais struct {
	ValorTotalNF float64 `json:"valor_total_nf"`
	ValorICMS    float64 `json:"valor_icms"`
}
