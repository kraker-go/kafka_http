package server

import (
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

func NewServer(port string, logg *zap.Logger, rout *mux.Router) *http.Server {
	if err := http.ListenAndServe(port, rout); err != nil {
		return &http.Server{
			Addr:    port,
			Handler: rout,
		}

		logg.Info("Сервер запущен")
	}
	return nil
}
