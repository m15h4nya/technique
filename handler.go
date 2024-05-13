package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var _ Handler = Handle{}

type Handler interface {
	AllFiles(w http.ResponseWriter, r *http.Request)
	FileContent(w http.ResponseWriter, r *http.Request)
	EditFile(w http.ResponseWriter, r *http.Request)
}

type Handle struct {
	vault vaulter
}

func NewHandler() Handle {
	return Handle{
		vault: NewVault(),
	}
}

func writeMarshal(w http.ResponseWriter, data any, status int) {
	bytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		status = http.StatusInternalServerError
	}
	w.WriteHeader(status)
	_, err = w.Write(bytes)
	if err != nil {
		fmt.Println(err)
	}
}

// @Summary	All files existing in vault
// @Produce	json
// @Success	200	{array}	string
// @Failure	500
// @Router	/api/all_files [get]
func (h Handle) AllFiles(w http.ResponseWriter, r *http.Request) {
	content, err := h.vault.AllFiles()
	if err != nil {
		fmt.Println(err)
		writeMarshal(w, nil, http.StatusInternalServerError)
		return
	}

	writeMarshal(w, content, http.StatusOK)
}

// @Summary	File's content
// @Produce	plain
// @Success	200
// @Failure	500
// @Param	filename	path	string	true	"Filename"
// @Router	/api/file/{filename} [get]
func (h Handle) FileContent(w http.ResponseWriter, r *http.Request) {
	filename := mux.Vars(r)["filename"]

	content, err := h.vault.FileContent(filename)
	if err != nil {
		fmt.Println(err)
		writeMarshal(w, nil, http.StatusInternalServerError)
		return
	}

	writeMarshal(w, content, http.StatusOK)
}

func (h Handle) EditFile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Edit is not possible yet")
}
