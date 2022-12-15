package models

import "time"

/* GraoTweet es el formato o estructura que tendrá nuestro Tweet */
type GraboTweet struct {
	/* bson corresponde al nombre con el que se guarda en BD y json la representación en json cuando se devuelva */
	UserID  string    `bson:"userid" json:"userid,omitempty"`
	Mensaje string    `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
}
