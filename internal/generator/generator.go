// internal/generator/generator.go

package generator

import (
	"math/rand"
	"time"

	"github.com/danrodsg/gerador-nfe-go/internal/models"
)

func gerarChaveAcesso() string {

	return "12345678901234567890123456789012345678901234"
}

func gerarCNPJFake() string {

	return "99.999.999/0001-99"
}

func GerarNotaFiscal() models.NotaFiscal {
	rand.Seed(time.Now().UnixNano())

	itens := gerarItensAleatorios()
	var totalItens float64
	for _, item := range itens {
		totalItens += item.ValorUnitario * float64(item.Quantidade)
	}

	nf := models.NotaFiscal{
		ChaveAcesso: gerarChaveAcesso(),
		DataEmissao: time.Now(),
		Emitente: models.Emitente{
			CNPJ: gerarCNPJFake(),
			Nome: "EMPRESA TESTE LTDA",
		},
		Destinatario: models.Destinatario{
			CPF_CNPJ: "111.111.111-11",
			Nome:     "Cliente de Teste",
		},
		Itens: itens,
		Totais: models.Totais{
			ValorTotalNF: totalItens * 1.18,
			ValorICMS:    totalItens * 0.18,
		},
	}
	return nf
}

func gerarItensAleatorios() []models.Item {
	numItens := rand.Intn(5) + 1
	items := make([]models.Item, numItens)

	return items
}
