package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

func NewServer(port string, logg *zap.Logger, rout *mux.Router) error {
	if err := http.ListenAndServe(port, rout); err != nil {
		return fmt.Errorf("ошибка: сервер не запущен %w", err)
	}

	logg.Info("Сервер запущен")

	return nil
}
