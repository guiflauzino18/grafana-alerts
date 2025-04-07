package controller

import (
	"fmt"
	"grafan-alerts/src/response"
	"grafan-alerts/src/utils"
	"net/http"
)

func DisparaAlerta(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Alerta recebido")

	Play("C:\\Grafana Alerts\\alert.wav")

	response.JSON(w, http.StatusOK, response.Sucesso{Retorno: "Alerta processado"})

}

func Play(audio string) {

	go func() {
		if erro := utils.PlayAudio(audio); erro != nil {
			fmt.Errorf("Erro ao reproduzir áudio:\n %v", erro)
		}
	}()

}
