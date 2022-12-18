package models

/*RespuestaConsultaRelacion tiene el true o false que se obtiende de consultar la relación
entre dos usuarios*/
type RespuestaConsultaRelacion struct {
	Status bool `json:"status"`
}
