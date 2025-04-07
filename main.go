package main

import (
	"fmt"
	"grafan-alerts/src/service"
	"log"

	"golang.org/x/sys/windows/svc"
)

func main() {
	//service.StartWebServer()

	isService, erro := service.IsWindowsService()
	if erro != nil {
		log.Fatalf("Erro ao verificar serviço:\n%v", erro)
	}

	if isService {
		erro = svc.Run(service.ServiceName, &service.MyService{})
		if erro != nil {
			log.Fatalf("Erro ao executar serviço:\n%v", erro)
		}
	} else {
		fmt.Println("Rodando em modo console")
		service.StartWebServer()
	}
}
