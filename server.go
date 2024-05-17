package main

import (
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/m15h4nya/online_vault/docs"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// @title			Technique
// @version		1.0
// @description	vault reader
func NewServer() http.Server {
	return http.Server{
		Handler: router(NewHandler()),
		Addr:    "0.0.0.0:8080",
	}
}

func router(handler Handle) *mux.Router {
	r := mux.NewRouter()
	r.Use(CorsMiddlware)

	getHandler := r.Methods(http.MethodGet).Subrouter()
	getHandler.HandleFunc("/api/all_files", handler.AllFiles)
	getHandler.HandleFunc("/api/file/{filename}", handler.FileContent)

	postHandler := r.Methods(http.MethodPost).Subrouter()
	postHandler.HandleFunc("/api/file/{filename}", handler.EditFile)

	getHandler.PathPrefix("/swagger/").HandlerFunc(httpSwagger.Handler())

	return r
}
