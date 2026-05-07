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

func (h *MovieHandler) DeleteMovieByID(w http.ResponseWriter, r *http.Request) {
	h.Logg.Info("удаляем видео файлов")
	ctx := r.Context()
	movieID := mux.Vars(r)["id"]
	id, err := strconv.Atoi(movieID)
	if err != nil {
		http.Error(w, "ошибка удаления фильма", http.StatusBadRequest)
		h.Logg.Error("handler: удаления фильма")
		return
	}

	err = h.Hand.DeleteMovie(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrorIdIsEmpty) {
			http.Error(w, "ID не может быть ноль и меньше", http.StatusBadRequest)
			h.Logg.Error("ввели ноль или отрицательный id")
			return
		}
		if errors.Is(err, domain.ErrorMovieNotFound) {
			http.Error(w, "фильм с таким ID не найден", http.StatusBadRequest)
			h.Logg.Error("фильм с таким ID не найден")
			return
		}
		http.Error(w, "ошибка удаления фильма", http.StatusBadRequest)
		h.Logg.Error("handler: ошибка удаления фильма", zap.Error(err))
		return
	}

	h.Logg.Info("уделен фаил с ", zap.Int("ID:", id))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"удален фильм с ID:": id})
}
