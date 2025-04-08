package service

import (
	"fmt"
	"grafan-alerts/src/rotas"
	"log"
	"net/http"
	"os"
	"os/exec"

	"golang.org/x/sys/windows/svc"
)

const ServiceName = "GrafanaAlerts"
const DisplayName = "Grafana Alertas"

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
				s <- svc.Status{State: svc.StopPending}

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
	isService, erro := svc.IsWindowsService()
	return isService, erro

}

func InstallService() {
	CaminhoExe, erro := os.Executable()
	if erro != nil {
		log.Fatalf("Erro ao abri caminho do executavel:\n%w", erro)
	}

	cmd := exec.Command("sc", "create", ServiceName, "binPath=", fmt.Sprintf("\"%s\"", CaminhoExe),
		"DisplayName=", DisplayName, "start=", "auto")

	output, erro := cmd.CombinedOutput()
	if erro != nil {
		log.Fatalf("Erro ao instalar serviço:\n%v", erro, string(output))
	}

	log.Println("Serviço instalado")
}

func UninstallService() {
	cmd := exec.Command("sc", "delete", ServiceName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Erro ao remover serviço: %v\n%s", err, string(output))
	}
	log.Println("Serviço removido com sucesso!")
}
