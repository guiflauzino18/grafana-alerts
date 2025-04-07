package service

import (
	"fmt"
	"grafan-alerts/src/rotas"
	"log"
	"net/http"

	"golang.org/x/sys/windows/svc"
)

const ServiceName = "GrafanaAlerts"

type MyService struct{}

func (m *MyService) Execute(args []string, r <-chan svc.ChangeRequest, s chan<- svc.Status) (bool, uint32) {

	s <- svc.Status{State: svc.StartPending}

	//Inicia servidor web aqui
	go StartWebServer()

	s <- svc.Status{State: svc.Running, Accepts: svc.AcceptStop | svc.AcceptShutdown}

	for {
		select {
		case c := <-r:
			switch c.Cmd {
			case svc.Stop, svc.Shutdown:
				s <- svc.Status{State: svc.Shutdown}

				return false, 0
			}
		}
	}
}

func StartWebServer() {
	r := rotas.GerarRotas()

	porta := 5000

	fmt.Printf("Aplicação executando na porta %d", porta)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", porta), r))
}

func IsWindowsService() (bool, error) {
	isService, erro := svc.isWindowsService
	return isService, erro

}
