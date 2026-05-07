package handler

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"kafka_http/internal/domain"
	"net/http"
	"strconv"
)

func (h *MovieHandler) GetMovieByID(w http.ResponseWriter, r *http.Request) {
	h.Logg.Info("поиск видео по ID")
	ctx := r.Context()
	movieID := mux.Vars(r)["id"]
	id, err := strconv.Atoi(movieID)
	if err != nil {
		http.Error(w, "неверный формат ID", http.StatusBadRequest)
		h.Logg.Info("неверный формат ID")
		return
	}

	movie, err := h.Hand.GetMovieByID(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrorIdIsEmpty) {
			http.Error(w, "ID не может быть ноль и меньше", http.StatusBadRequest)
			h.Logg.Error("ввели ноль или отрицательный id")
			return
		}
		http.Error(w, "ошибка поиска фильма", http.StatusBadRequest)
		h.Logg.Error("handler: поиска фильма", zap.Error(err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movie)

}
