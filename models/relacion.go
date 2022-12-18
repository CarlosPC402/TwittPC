package models

/* Relacion modelo para grabar la relación de un usuario con otro, el primero corresponde al usuario
mismo y el segundo al cual se va relacionar(usuario al que va seguir) */
type Relacion struct {
	UsuarioID         string `bson:"usuarioid" json:"usuarioId"`
	UsuarioRelacionID string `bson:"usuariorelacionid" json:"usuarioRelacionId"`
}
