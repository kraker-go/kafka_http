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

func (h *MovieHandler) UpdateMovieByMovieID(w http.ResponseWriter, r *http.Request) {
	h.Logg.Info("обновляем данные")
	ctx := r.Context()
	movieID := mux.Vars(r)["id"]
	id, err := strconv.Atoi(movieID)
	if err != nil {
		http.Error(w, "неверный формат ID", http.StatusBadRequest)
		h.Logg.Error("ввели неверный ID")
		return
	}

	var update *domain.Movie
	if err = json.NewDecoder(r.Body).Decode(&update); err != nil {
		http.Error(w, "ошибка чтения JSON", http.StatusBadRequest)
		h.Logg.Error("ошибка чтения JSON", zap.Error(err))
		return
	}

	movie, err := h.Hand.UpdateMovie(ctx, id, update)
	if err != nil {
		if errors.Is(err, domain.ErrorMovieNotFound) {
			http.Error(w, "фаил с таким ID не найден", http.StatusBadRequest)
			h.Logg.Error("фаил с таким ID не найден")
			return
		}
		if errors.Is(err, domain.ErrorIdIsEmpty) {
			http.Error(w, "ID не может быть ноль и меньше", http.StatusBadRequest)
			h.Logg.Error("ввели ноль или отрицательный id")
			return
		}

		http.Error(w, "ошибка обновления фильма", http.StatusBadRequest)
		h.Logg.Error("handler: ошибка обновления фильма", zap.Error(err))
		return
	}

	h.Logg.Info("данные обновлены!")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"данные обновлены: ": movie,
	})

}
