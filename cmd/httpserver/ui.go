package httpserver

import (
	"html/template"
	"log"
	"net/http"
	"path"

	"github.com/julienschmidt/httprouter"
)

func HtmlTemplate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	// Parse file
	parsedTemplate, _ := template.ParseFiles(path.Join("public", "index.html"))
	// Parse template
	err := parsedTemplate.Execute(w, "")
	if err != nil {
		log.Printf("Error executing template: %v", err)
		return
	}

}
