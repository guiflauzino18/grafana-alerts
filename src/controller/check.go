package controller

import (
	"fmt"
	"grafan-alerts/src/response"
	"net/http"
)

func Chech(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("Requisição recebida pelo ip: %s", r.Host)
	response.JSON(w, http.StatusOK, response.Sucesso{Retorno: "Up"})
}
