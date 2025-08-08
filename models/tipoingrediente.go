package models

type TipoIngrediente struct {
	COD          int      `json:"codtipoingrediente"`
	Nombre       string   `json:"nombreti"`
	UnidadMedida []string `json:"unidadmedida"`
}
