package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewServer() http.Server {
	return http.Server{
		Handler: router(NewHandler()),
		Addr:    "127.0.0.1:8080",
	}
}

func router(handler Handler) *mux.Router {
	r := mux.NewRouter()

	getHandler := r.Methods(http.MethodGet).Subrouter()
	getHandler.HandleFunc("/api/all_files", handler.AllFiles)
	getHandler.HandleFunc("/api/file/{filename}", handler.FileContent)
	getHandler.HandleFunc("/api/file/{filename}/links", handler.FileLinks)

	postHandler := r.Methods(http.MethodPost).Subrouter()
	postHandler.HandleFunc("/api/file/{filename}", handler.EditFile)

	return r
}
