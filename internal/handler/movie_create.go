package handler

import (
	"encoding/json"
	"errors"
	"go.uber.org/zap"
	"kafka_http/internal/domain"
	"net/http"
)

func (h *MovieHandler) CreateMovie(w http.ResponseWriter, r *http.Request) {
	h.Logg.Info("создаем видео фаил")
	ctx := r.Context()

	var mov domain.Movie

	if err := json.NewDecoder(r.Body).Decode(&mov); err != nil {
		http.Error(w, "handler: ошибка json", http.StatusBadRequest)
		h.Logg.Error("handler: ошибка json", zap.Error(err))
		return
	}

	err := h.Hand.CreateMovie(ctx, &mov)
	if err != nil {
		if errors.Is(err, domain.ErrorMovieExist) {
			http.Error(w, "фаил с таким названием уже в базе", http.StatusBadRequest)
			h.Logg.Error("такой фаил уже есть в базе данных")
			return
		}
		http.Error(w, "ошибка создания фильма", http.StatusBadRequest)
		h.Logg.Error("handler: ошибка создания фильма", zap.Error(err))
		return
	}

	h.Logg.Info("фильм успешно добавлен")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(mov); err != nil {
		http.Error(w, "handler: ошибка ответа json", http.StatusBadRequest)
		h.Logg.Error("handler: ошибка ответа json", zap.Error(err))
		return
	}
}
