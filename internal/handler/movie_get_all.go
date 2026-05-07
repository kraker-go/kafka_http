package handler

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
)

func (h *MovieHandler) GetMovies(w http.ResponseWriter, r *http.Request) {
	h.Logg.Info("поиск всех видео файлов")

	ctx := r.Context()

	movies, err := h.Hand.GetMovies(ctx)
	if err != nil {
		http.Error(w, "ошибка поиска фильмов", http.StatusBadRequest)
		h.Logg.Error("handler: поиска фильмов", zap.Error(err))
		return
	}

	h.Logg.Info("найдено фильмов - ", zap.Int("", len(movies)))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movies)
}
