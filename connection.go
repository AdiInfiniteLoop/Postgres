package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

const (
	successStatusCode = http.StatusOK
	successMessage    = "Hello to the Server"
)

func Connect() {
	router := chi.NewRouter()

	srv := &http.Server{
		Handler: router,
		Addr:    ":8080",
	}

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(successStatusCode)
		_, err := w.Write([]byte(successMessage))
		if err != nil {
			return
		}
	})
	fmt.Println("Server running...")
	err := srv.ListenAndServe()
	if err != nil {
		log.Println("Error Occurred", err)
	}
}
