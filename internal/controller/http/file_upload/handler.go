package file_upload

import (
	"card-detect-demo/internal/model"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Boxes []model.Box `json:"boxes"`
	Img   string      `json:"img"`
}

type Detector interface {
	Detect(pathImg string) ([]model.Box, string, error)
}

type FileUploadHandler struct {
	name       string
	detector   Detector
	dirTmpPath string
}

func NewFileUploadHandler(detector Detector, dirTmpPath string) *FileUploadHandler {
	return &FileUploadHandler{
		name:       "FileUploadHandler",
		detector:   detector,
		dirTmpPath: dirTmpPath,
	}
}

func (h *FileUploadHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseMultipartForm(10 << 20) // 10MB max
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get file from form data
	file, handler, err := r.FormFile("image")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
	defer file.Close()

	// Save the file to disk
	fileName := h.dirTmpPath + "/" + handler.Filename
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()

	io.Copy(f, file)

	boxes, path, err := h.detector.Detect(fileName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	resp := Response{
		Boxes: boxes,
		Img:   path,
	}
	json.NewEncoder(w).Encode(resp)
}
