package main

import (
	"log"
	"net/http"

	"github.com/iglesiasgm/wiki-cocktail/db"
	"github.com/iglesiasgm/wiki-cocktail/handlers"
)

func main() {
	db.InitDB()
	http.HandleFunc("/garnish", handlers.GetAllGarnishes)
	log.Println("Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

/*

PROBAR ESCRIBIENDO ESTO EN EL BUSCADOR

http://localhost:8080/cocktails

http://localhost:8080/cocktails/1

*/
