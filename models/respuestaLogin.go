package models

/* RespuestaLogin tiene el token que se devuelve en el Login */
type RespuestaLogin struct {
	Token string `json:"token,omitempty"`
}
