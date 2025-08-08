package models

type Cocktail struct {
	ID          int           `json:"id"`
	Nombre      string        `json:"nombrecoctel"`
	Description string        `json:"descripcoctel"`
	Vaso        string        `json:"idtipoVaso"`
	Hielo       string        `json:"idtipoHielo"`
	Canthielo   int           `json:"canthielo"`
	Ingredients []Ingredients `json:"ingredients"`
	Metodo      string        `json:"idmetodo"`
	Garnish     []Garnish     `json:"garnish"`
	Activo      bool          `json:"isactivecocktail"`
}
