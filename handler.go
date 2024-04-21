package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler interface {
	AllFiles(w http.ResponseWriter, r *http.Request)
	FileContent(w http.ResponseWriter, r *http.Request)
	FileLinks(w http.ResponseWriter, r *http.Request)
	EditFile(w http.ResponseWriter, r *http.Request)
}

type Handle struct {
	vault vaulter
}

func NewHandler() Handler {
	return Handle{
		vault: NewVault(),
	}
}

func (h Handle) AllFiles(w http.ResponseWriter, r *http.Request) {
	content, err := h.vault.AllFiles()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
	}

	fmt.Fprint(w, content)
}

func (h Handle) FileContent(w http.ResponseWriter, r *http.Request) {
	filename := mux.Vars(r)["filename"]

	content, err := h.vault.FileContent(filename)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
	}

	fmt.Fprint(w, content)
}

func (h Handle) FileLinks(w http.ResponseWriter, r *http.Request) {
	panic("not implemented") // TODO: Implement
}

func (h Handle) EditFile(w http.ResponseWriter, r *http.Request) {
	panic("not implemented") // TODO: Implement
}
