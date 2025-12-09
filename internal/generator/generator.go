// internal/generator/generator.go

package generator

import (
    "math/rand"
    "time"
    // Importe o pacote de models que você criou
    "github.com/danrodsg/gerador-nfe-go/internal/models" 
)

// Gera uma string numérica de 44 dígitos
func gerarChaveAcesso() string {
    // Implementação da geração da chave (44 dígitos)
    return "12345678901234567890123456789012345678901234" // Exemplo de chave fixa
}

// gerarCNPJFake deve gerar um CNPJ válido/formatado para teste
func gerarCNPJFake() string {
    // Implementação da geração do CNPJ
    return "99.999.999/0001-99"
}

// GerarNotaFiscal orquestra a criação de uma NF-e completa.
func GerarNotaFiscal() models.NotaFiscal {
    rand.Seed(time.Now().UnixNano())
    
    // 1. Geração dos Itens e cálculo do subtotal
    itens := gerarItensAleatorios()
    var totalItens float64
    for _, item := range itens {
        totalItens += item.ValorUnitario * float64(item.Quantidade)
    }

    // 2. Montagem da NF-e
    nf := models.NotaFiscal{
        ChaveAcesso: gerarChaveAcesso(),
        DataEmissao: time.Now(),
        Emitente: models.Emitente{
            CNPJ: gerarCNPJFake(),
            Nome: "EMPRESA TESTE LTDA",
        },
        Destinatario: models.Destinatario{
            CPF_CNPJ: "111.111.111-11", // CPF fake
            Nome:     "Cliente de Teste",
        },
        Itens: itens,
        Totais: models.Totais{
            ValorTotalNF: totalItens * 1.18, // Exemplo: 18% de imposto
            ValorICMS:    totalItens * 0.18,
        },
    }
    return nf
}

// Função auxiliar (pode ficar no mesmo arquivo)
func gerarItensAleatorios() []models.Item {
    numItens := rand.Intn(5) + 1 // Gera entre 1 e 5 itens
    items := make([]models.Item, numItens)
    // Preenche a lista de items com dados fictícios
    // ...
    return items
}