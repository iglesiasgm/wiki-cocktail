package models

type Ingredients struct {
	ID        int      `json:"idingerdiente"`
	Receta    string   `json:"receta"`
	Tipo      []string `json:"tipoingrediente"`
	CantFinal int      `json:"cantfinal"`
	Activo    bool     `json:"isactiveingredient"`
}
