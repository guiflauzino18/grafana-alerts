package response

import (
	"encoding/json"
	"log"
	"net/http"
)

// Struct para criar um erro
type Erro struct {
	Erro string `json:"erro"`
}

type Sucesso struct {
	Retorno string `json:"retorno"`
}

func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if statusCode != http.StatusNoContent {
		if erro := json.NewEncoder(w).Encode(dados); erro != nil {
			log.Fatal(erro)
		}
	}
}

func TratarStatusDeErro(w http.ResponseWriter, r *http.Response) {
	var erro Erro
	json.NewDecoder(r.Body).Decode(&erro)
	JSON(w, r.StatusCode, erro)
}
