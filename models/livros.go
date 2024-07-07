package models

import  (
	
)

type Livros struct {
	Id        int
	Nome      string
	Descricao string
	Preco     float64
	Genero    string
}

func BuscaLivros() []Livros {

	db := Db_conect()

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

	defer.Close()
	return livros

}
