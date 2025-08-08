package models

type Cocktail struct {
	ID          int      `json:"id"`
	Nombre      string   `json:"nombrecoctel"`
	Description string   `json:"descripcoctel"`
	Vaso        string   `json:"idtipoVaso"`
	Hielo       string   `json:"idtipoHielo"`
	Ingredients []string `json:"ingredients"`
	Metodo      string   `json:"idmetodo"`
	Garnish     []string `json:"garnish"`
	Activo      bool     `json:"isactivecocktail"`
}
