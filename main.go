package main

import (
	"fmt"
	"grafan-alerts/src/service"
	"log"
	"os"

	"golang.org/x/sys/windows/svc"
)

func main() {

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "install":
			service.InstallService()
			return

		case "uninstall":
			service.UninstallService()
			return
		}
	}

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
