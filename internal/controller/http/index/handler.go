package index

import (
	"html/template"
	"log"
	"net/http"
)

type projectInfo struct {
	Name    string
	Version string
}

type Handler struct {
	name    string
	version string
}

func NewIndexHandler(name, version string) *Handler {
	return &Handler{
		name:    name,
		version: version,
	}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	data := projectInfo{
		Name:    h.name,
		Version: h.version,
	}

	// Парсим шаблон из файла
	tmpl, err := template.ParseFiles("./template/index.html")
	if err != nil {
		log.Fatal(err)
	}

	// Генерируем вывод на основе шаблона и данных
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
}
