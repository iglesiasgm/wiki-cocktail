package main

import (
	"log"
	"net/http"

	"github.com/iglesiasgm/wiki-cocktail/handlers"
)

func main() {
	http.HandleFunc("/cocktails", handlers.GetCocktails)
	http.HandleFunc("/cocktails/", handlers.GetCocktailByID)

	log.Println("Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

/* 

PROBAR ESCRIBIENDO ESTO EN EL BUSCADOR

http://localhost:8080/cocktails

http://localhost:8080/cocktails/1

*/