package main

import (
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {

	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))

	http.HandleFunc("/", index)

	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	db := db_conect()

	selectAllLivros, err := db.Query("select * from livros")

	if err != nil {
		panic(err.Error())
	}

	p := Livros{}
	livros := []Livros{}

	for selectAllLivros.Next() {
		var id int
		var nome string
		var descricao string
		var preco float64
		var genero string

		err = selectAllLivros.Scan(&id, &nome, &descricao, &preco, &genero)

		if err != nil {
			panic(err.Error())
		}

		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Genero = genero

		livros = append(livros, p)

	}

	temp.ExecuteTemplate(w, "index", livros)

	defer db.Close()
}
