package data

import "github.com/iglesiasgm/wiki-cocktail/models"

var Cocktails = []models.Cocktail{
	{ID: 1, Name: "Mojito", Ingredients: []string{"Rum", "Mint", "Lime", "Soda"}},
	{ID: 2, Name: "Margarita", Ingredients: []string{"Tequila", "Lime", "Triple sec"}},
	{ID: 3, Name: "Old Fashioned", Ingredients: []string{"Bourbon", "Angostura Bitters", "Sugar"}},
}
