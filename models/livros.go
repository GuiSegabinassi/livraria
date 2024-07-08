package models

import (
	"loja/db"
)

type Livros struct {
	Id        int
	Nome      string
	Descricao string
	Preco     float64
	Genero    string
}

func BuscaLivros() []Livros {

	db := db.Db_conect()

	// Lembre-se de fechar a conexão com o banco de dados
	defer db.Close()

	selectAllLivros, err := db.Query("select * from livros")
	if err != nil {
		panic(err.Error())
	}
	defer selectAllLivros.Close()

	livros := []Livros{}

	for selectAllLivros.Next() {
		var p Livros

		err = selectAllLivros.Scan(&p.Id, &p.Nome, &p.Descricao, &p.Preco, &p.Genero)
		if err != nil {
			panic(err.Error())
		}

		livros = append(livros, p)
	}

	return livros
}
