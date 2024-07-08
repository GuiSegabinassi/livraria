package main

import (
	"log"
	"loja/routers"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))

	routers.CarregaRotas()

	log.Println("Listening on :8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
