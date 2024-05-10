package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewServer() http.Server {
	return http.Server{
		Handler: router(NewHandler()),
		Addr:    "0.0.0.0:8080",
	}
}

func router(handler Handle) *mux.Router {
	r := mux.NewRouter()

	getHandler := r.Methods(http.MethodGet).Subrouter()
	getHandler.HandleFunc("/api/all_files", handler.AllFiles)
	getHandler.HandleFunc("/api/file/{filename}", handler.FileContent)

	postHandler := r.Methods(http.MethodPost).Subrouter()
	postHandler.HandleFunc("/api/file/{filename}", handler.EditFile)

	return r
}
