package handlers

import (
	"io"
	"net/http"
	"os"
	"time"

	"github.com/Cornpop456/go-6-sprint-final/internal/service"
)

const (
	indexPage = "index.html"
	keyFile   = "myFile"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	file, err := os.Open(indexPage)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer file.Close()

	data, err := io.ReadAll(file)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func HandleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(32 << 20)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	file, _, err := r.FormFile(keyFile)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer file.Close()

	data, err := io.ReadAll(file)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	outFile, err := os.OpenFile(time.Now().UTC().String()+".txt", os.O_WRONLY|os.O_CREATE, 0755)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer outFile.Close()

	converted := service.Convert(string(data))

	_, err = outFile.WriteString(converted)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(converted))
}
