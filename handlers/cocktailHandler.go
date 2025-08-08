package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/iglesiasgm/wiki-cocktail/db"
	"github.com/iglesiasgm/wiki-cocktail/models"
)

func GetAllCocktails(w http.ResponseWriter, r *http.Request) {

	queryCocteles := `
		SELECT 
			c.idcoctel, c.nombrecoctel, c.descripcoctel, c.canthielo, c.isactivecoctail,
			m.descripmetodo,
			th.nombretipohielo,
			tv.nombretipovaso
		FROM coctel c
		LEFT JOIN metodo m ON c.idmetodo = m.idmetodo
		LEFT JOIN tipohielo th ON c.idtipohielo = th.idtipohielo
		LEFT JOIN tipovaso tv ON c.idtipovaso = tv.idtipovaso
	`
	rows, err := db.DB.Query(queryCocteles)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	coctelesMap := make(map[int]*models.Cocktail)

	for rows.Next() {
		var c models.Cocktail
		if err := rows.Scan(
			&c.ID, &c.Nombre, &c.Description, &c.Canthielo, &c.Activo,
			&c.Metodo, &c.Hielo, &c.Vaso,
		); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		c.Garnish = []models.Garnish{}
		c.Ingredients = []models.Ingredients{}
		coctelesMap[c.ID] = &c
	}

	queryGarnish := `
		SELECT gc.idcoctel, g.idgarnish, g.descripgarnish
		FROM garnishcoctel gc
		LEFT JOIN garnish g ON gc.idgarnish = g.idgarnish
	`
	rowsG, err := db.DB.Query(queryGarnish)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rowsG.Close()

	for rowsG.Next() {
		var idCoctel int
		var g models.Garnish
		if err := rowsG.Scan(&idCoctel, &g.ID, &g.Description); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if c, ok := coctelesMap[idCoctel]; ok {
			c.Garnish = append(c.Garnish, g)
		}
	}

	queryIngredientes := `
		SELECT ic.idcoctel, ing.idingrediente, ing.receta, ing.cantfinal,
		       ti.nombretit, um.nomunidadmedida
		FROM ingredientecoctel ic
		LEFT JOIN ingrediente ing ON ic.idingrediente = ing.idingrediente
		LEFT JOIN tipoingrediente ti ON ing.codtipoingrediente = ti.codtipoingrediente
		LEFT JOIN tipoingredienteunidadmedida tiu ON ti.codtipoingrediente = tiu.codtipoingrediente
		LEFT JOIN unidadmedida um ON tiu.codum = um.codum
	`
	rowsI, err := db.DB.Query(queryIngredientes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rowsI.Close()

	for rowsI.Next() {
		var idCoctel int
		var i models.Ingredients
		if err := rowsI.Scan(
			&idCoctel, &i.ID, &i.Receta, &i.CantFinal, &i.Tipo,
		); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if c, ok := coctelesMap[idCoctel]; ok {
			c.Ingredients = append(c.Ingredients, i)
		}
	}

	cocteles := make([]models.Cocktail, 0, len(coctelesMap))
	for _, c := range coctelesMap {
		cocteles = append(cocteles, *c)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cocteles)
}
