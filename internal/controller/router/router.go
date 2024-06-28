package router

import (
	"card-detect-demo/internal/controller/http/file_upload"
	"card-detect-demo/internal/controller/http/index"
	"net/http"
)

func NewRouter(detectService file_upload.Detector, tmpFilePath, name, version string) *http.ServeMux {
	mux := http.NewServeMux()

	// Создаем файловый сервер, который будет использовать директорию "./template/static"
	staticStorage := http.FileServer(http.Dir("./template/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", staticStorage))
	fileStorage := http.FileServer(http.Dir("./storage"))
	mux.Handle("/storage/", http.StripPrefix("/storage/", fileStorage))

	indexHandler := index.NewIndexHandler(name, version)
	uploadHandler := file_upload.NewFileUploadHandler(detectService, tmpFilePath)

	mux.HandleFunc("/", indexHandler.Handle)
	mux.HandleFunc("/detect", uploadHandler.Handle)

	return mux
}
