// cmd/api/main.go

package main

import (
    "encoding/json"
    "log"
    "net/http"
    // Importe a função de geração
    "github.com/danrodsg/gerador-nfe-go/internal/generator" 
)

func handlerGerarNFe(w http.ResponseWriter, r *http.Request) {
    // 1. Chamar a função principal de geração
    nfFake := generator.GerarNotaFiscal()

    // 2. Configurar o cabeçalho para JSON
    w.Header().Set("Content-Type", "application/json")
    
    // 3. Serializar a struct para JSON e enviar
    json.NewEncoder(w).Encode(nfFake)
}

func main() {
    // Define o endpoint e qual função ele chama
    http.HandleFunc("/api/v1/nfe/generate", handlerGerarNFe) 

    log.Println("Servidor iniciado em http://localhost:8080")
    // Inicia o servidor na porta 8080
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Erro ao iniciar o servidor: %v", err)
    }
}