package models

type UnidadMedida struct {
	COD      int    `json:"codUM"`
	UM       string `json:"unidadMedida"`
	CANTIDAD int    `json:"cantidadIC"`
}
