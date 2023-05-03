package controller

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"main/internal/models"
	"net/http"
)

func RunHttp(addr string, dbi models.DBinteraction) {
	router := chi.NewRouter()

	router.Post("/user", func(w http.ResponseWriter, r *http.Request) {
		CreateHandler(w, r, dbi)
	})
	router.Post("/user/{user_id}/friends", func(w http.ResponseWriter, r *http.Request) {
		MakeFriendsHandler(w, r, dbi)
	})
	router.Delete("/user/{user_id}", func(w http.ResponseWriter, r *http.Request) {
		DelUserHandler(w, r, dbi)
	})
	router.Get("/user/{user_id}/friends", func(w http.ResponseWriter, r *http.Request) {
		GetFriendsHandler(w, r, dbi)
	})
	router.Get("/users", func(w http.ResponseWriter, r *http.Request) {
		GetUsersHandler(w, r, dbi)
	})
	router.Put("/user/{user_id}/age", func(w http.ResponseWriter, r *http.Request) {
		AgeUpdateHandler(w, r, dbi)
	})
	fmt.Println("Server started")
	err := http.ListenAndServe(addr, router)
	if err != nil {
		log.Fatalln(err)
	}

}
