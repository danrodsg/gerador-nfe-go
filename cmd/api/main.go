package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/danrodsg/gerador-nfe-go/internal/generator"
)

func handlerGerarNFe(w http.ResponseWriter, r *http.Request) {

	nfFake := generator.GerarNotaFiscal()

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(nfFake)
}

func main() {

	http.HandleFunc("/api/v1/nfe/generate", handlerGerarNFe)

	log.Println("Servidor iniciado em http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
