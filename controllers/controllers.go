package controllers

import (
	"html/template"
	"loja/models"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allLivros := models.BuscaLivros()
	temp.ExecuteTemplate(w, "index", allLivros)
}
