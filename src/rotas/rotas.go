package rotas

import (
	"grafan-alerts/src/controller"
	"net/http"

	"github.com/gorilla/mux"
)

// Struct de Rotas
type Rota struct {
	URI    string
	Metodo string
	Funcao func(http.ResponseWriter, *http.Request)
}

var rotasWebhook = []Rota{
	{
		URI:    "/alerta",
		Metodo: http.MethodPost,
		Funcao: controller.DisparaAlerta,
	},
	{
		URI:    "/check",
		Metodo: http.MethodGet,
		Funcao: controller.Chech,
	},
}

// Configura as rotas
func ConfigurarRotas(router *mux.Router) *mux.Router {
	rotas := rotasWebhook

	for _, rota := range rotas {
		//router.HandlerFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
		router.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}

	return router
}

func GerarRotas() *mux.Router {
	r := mux.NewRouter()
	return ConfigurarRotas(r)
}
