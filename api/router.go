package api

import (
	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"ramazon/service"
)

type routes struct {
	root    *mux.Router
	apiRoot *mux.Router
}

type api struct {
	routes         *routes
	ramazonService service.RamazonService
	logger         *zap.Logger
}

func Init(
	root *mux.Router,
	ramazon service.RamazonService,
	logger *zap.Logger) {

	r := routes{
		root:    root,
		apiRoot: root.PathPrefix("/api").Subrouter(),
	}

	api := api{
		routes:         &r,
		ramazonService: ramazon,
	}

	api.initRamazon()
}

func (api *api) initRamazon() {
	api.routes.apiRoot.HandleFunc("/pray-time/set", api.SetPrayTime).Methods("POST")
}
