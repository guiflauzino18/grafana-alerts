package controller

import (
	"fmt"
	"grafan-alerts/src/response"
	"grafan-alerts/src/utils"
	"net/http"
)

func DisparaAlerta(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Alerta recebido")

	if erro := Play("C:\\Grafana Alerts\\alert.wav"); erro != nil {
		response.JSON(w, http.StatusOK, response.Erro{Erro: erro.Error()})
		return
	}

	response.JSON(w, http.StatusOK, response.Sucesso{Retorno: "Alerta processado"})
}

func Play(audio string) error {

	erroChan := make(chan error)

	go func(canal chan error) {
		if erro := utils.PlayAudio(audio); erro != nil {
			canal <- fmt.Errorf(fmt.Sprintf("Erro ao reproduzir áudio:%v", erro))
		}

		close(canal)

	}(erroChan)

	var erros error

	for erro := range erroChan {
		erros = erro
	}

	return erros

}
