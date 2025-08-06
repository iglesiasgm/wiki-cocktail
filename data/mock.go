package data

import "github.com/iglesiasgm/wiki-cocktail/models"

var Cocktails = []models.Cocktail{
	{ID: 1, Name: "Mojito", Ingredients: []string{"Rum", "Mint", "Lime", "Soda"}, Garnish: "Mint leaves"},
	{ID: 2, Name: "Margarita", Ingredients: []string{"Tequila", "Lime", "Triple sec"}, Garnish: "Lemon Slice"},
	{ID: 3, Name: "Old Fashioned", Ingredients: []string{"Bourbon", "Angostura Bitters", "Sugar"}, Garnish: "Orange Peel"},
}
