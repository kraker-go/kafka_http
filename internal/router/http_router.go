package router

import (
	"github.com/gorilla/mux"
	"kafka_http/internal/handler"
)

func NewRouter(hand handler.MovieHandler) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/movie", hand.CreateMovie).Methods("POST")
	router.HandleFunc("/movie", hand.GetMovies).Methods("GET")
	router.HandleFunc("/movie/{id}", hand.GetMovieByID).Methods("GET")
	router.HandleFunc("/movie/{id}", hand.DeleteMovieByID).Methods("DELETE")
	router.HandleFunc("/movie/{id}", hand.UpdateMovieByMovieID).Methods("PUT")

	return router
}
