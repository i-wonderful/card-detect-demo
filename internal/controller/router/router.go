package router

import (
	"card-detect-demo/internal/controller/http/file_upload"
	"card-detect-demo/internal/controller/http/index"
	"net/http"
)

func NewRouter(detectService file_upload.Detector, tmpFilePath, version string) *http.ServeMux {
	// Создаем файловый сервер, который будет использовать директорию "./template/static"
	fs := http.FileServer(http.Dir("./template/static"))

	// Создаем ServeMux, который будет маршрутизировать запросы
	mux := http.NewServeMux()

	// Правильно указываем паттерн и убираем префикс
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	fsStorage := http.FileServer(http.Dir("./storage"))
	mux.Handle("/storage/", http.StripPrefix("/storage/", fsStorage))

	indexHandler := index.NewIndexHandler(version)
	uploadHandler := file_upload.NewFileUploadHandler(detectService, tmpFilePath)

	mux.HandleFunc("/", indexHandler.Handle)
	mux.HandleFunc("/detect", uploadHandler.Handle)

	return mux
}
